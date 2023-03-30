package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func download(url string, threads int) {
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

	contentBody := make([]string, threads+1)

	for i := 0; i < threads; i++ {
		wg.Add(1)

		min := individualLength * i
		max := individualLength * (i + 1)

		if i == threads-1 {
			max += remainder
		}

		go func(min int, max int, i int) {
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

			reader, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			contentBody[i] = string(reader)
			err = os.WriteFile(strconv.Itoa(i), []byte(contentBody[i]), 0x777)
			if err != nil {
				return
			}

			defer wg.Done()
		}(min, max, i)
	}

	wg.Wait()
}
