/*
Copyright © 2023 simonfalke ziqibrandonli@gmail.com
*/

package downloader

import (
	"fmt"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Downloader struct {
	Url      string
	Threads  int
	SavePath string
}

type safeFileContent struct {
	mu sync.Mutex
	v  map[int]string
}

func (f *safeFileContent) Set(part int, content string) {
	f.mu.Lock()
	f.v[part] = content
	f.mu.Unlock()
}

func (f *safeFileContent) Value(part int) string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.v[part]
}

var wg sync.WaitGroup
var wgBar sync.WaitGroup

func Download(d Downloader) {
	url := d.Url
	threads := d.Threads
	savePath := d.SavePath

	defer func() {
		if r := recover(); r != nil {
			log.Fatal("A fatal error has occurred: ", r)
		}
	}()

	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}

	maps := res.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])
	if err != nil {
		panic(err)
	}

	individualLength := length / threads
	remainder := length % threads

	p := mpb.New(mpb.WithWaitGroup(&wgBar))
	numBars := threads
	wgBar.Add(numBars)

	fileContent := safeFileContent{v: make(map[int]string)}

	for i := 0; i < threads; i++ {
		wg.Add(1)

		min := individualLength * i
		max := individualLength * (i + 1)

		if i == threads-1 {
			max += remainder + 1
		}

		name := fmt.Sprintf("Thread %d ", i)
		//name := fmt.Sprintf("Thread#%d:", i)
		bar := p.AddBar(int64(individualLength),
			mpb.PrependDecorators(
				decor.Name(name),
				decor.Counters(decor.UnitKiB, "%.2f / %.2f"),
				decor.Percentage(decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				decor.EwmaETA(decor.ET_STYLE_GO, 30),
				decor.Name(" ] "),
				decor.EwmaSpeed(decor.UnitKiB, "%.2f", 30),
			),
		)

		go func(min int, max int, i int) {
			defer func() {
				if r := recover(); r != nil {
					log.Fatal("A fatal error has occurred: ", r)
				}
			}()

			//fmt.Println("Downloading part " + strconv.Itoa(i) + " from " + strconv.Itoa(min) + " to " + strconv.Itoa(max))

			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				panic(err)
			}

			rangeHeader := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1)
			req.Header.Add("Range", rangeHeader)
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(resp.Body)

			proxyReader := bar.ProxyReader(resp.Body)

			defer func(proxyReader io.ReadCloser) {
				err := proxyReader.Close()
				if err != nil {
					panic(err)
				}
			}(proxyReader)

			reader, err := io.ReadAll(proxyReader)
			if err != nil {
				panic(err)
			}

			fileContent.Set(i, string(reader))

			defer wg.Done()
			defer wgBar.Done()
		}(min, max, i)
	}

	wg.Wait()
	p.Wait()

	f, err := os.OpenFile(savePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	if err := f.Chmod(0666); err != nil {
		panic(err)
	}

	for i := 0; i < threads; i++ {
		content := fileContent.Value(i)

		if _, err := f.WriteString(content); err != nil {
			panic(err)
		}
		if err := f.Chmod(0666); err != nil {
			panic(err)
		}
	}

	fmt.Println("Download complete!")
}
