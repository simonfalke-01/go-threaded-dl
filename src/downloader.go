package main

import (
	"fmt"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var wgBar sync.WaitGroup

func download(d Downloader) {
	url := d.url
	threads := d.threads
	savePath := d.savePath

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

			contentBody := string(reader)
			err = os.WriteFile(filepath.Join("/tmp", "part"+strconv.Itoa(i)), []byte(contentBody), 0666)
			if err != nil {
				panic(err)
			}

			defer wg.Done()
			defer wgBar.Done()
		}(min, max, i)
	}

	wg.Wait()
	p.Wait()

	if stat, err := os.Stat(filepath.Dir(savePath)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(savePath), 0666); err != nil {
			panic(err)
		}
	} else if !stat.IsDir() {
		panic("Save path's parent is not a directory")
	}

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
		fTmp, err := os.OpenFile(filepath.Join("/tmp", "part"+strconv.Itoa(i)), os.O_RDONLY, 0666)
		if err != nil {
			panic(err)
		}

		content, err := io.ReadAll(fTmp)
		if err != nil {
			panic(err)
		}

		err = fTmp.Close()
		if err != nil {
			panic(err)
		}

		err = os.Remove(filepath.Join("/tmp", "part"+strconv.Itoa(i)))
		if err != nil {
			panic(err)
		}

		if _, err := f.Write(content); err != nil {
			panic(err)
		}
		if err := f.Chmod(0666); err != nil {
			panic(err)
		}
	}

	fmt.Println("Download complete!")
}
