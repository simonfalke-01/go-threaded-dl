package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func download(url string, threads int, savePath string) {
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

	for i := 0; i < threads; i++ {
		wg.Add(1)

		min := individualLength * i
		max := individualLength * (i + 1)

		if i == threads-1 {
			max += remainder + 1
		}

		go func(min int, max int, i int) {
			defer func() {
				if r := recover(); r != nil {
					log.Fatal("A fatal error has occurred: ", r)
				}
			}()

			fmt.Println("Downloading part " + strconv.Itoa(i) + " from " + strconv.Itoa(min) + " to " + strconv.Itoa(max))

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

			fmt.Println("Thread " + strconv.Itoa(i) + "status: " + resp.Status + " " + resp.Proto)

			reader, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			contentBody := string(reader)
			err = os.WriteFile(strconv.Itoa(i), []byte(contentBody), 0x777)
			if err != nil {
				panic(err)
			}

			f, err := os.OpenFile(savePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0x777)
			if err != nil {
				panic(err)
			}

			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					panic(err)
				}
			}(f)

			if _, err := f.WriteString(contentBody); err != nil {
				panic(err)
			}

			defer wg.Done()
		}(min, max, i)
	}

	wg.Wait()
}
