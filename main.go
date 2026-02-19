package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type Data struct {
	d  map[string]string
	mu sync.Mutex
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	urls := os.Args[1:]

	var waitter sync.WaitGroup
	fetchedData := Data{make(map[string]string), sync.Mutex{}}
	for _, url := range urls {
		waitter.Add(1)
		go FetchData(url, &fetchedData, &waitter)
	}

	waitter.Wait()
	for _, i := range fetchedData.d {
		fmt.Print("\n-------------------\n", i, "\n-------------------\n")
	}
}

func FetchData(url string, fetchedData *Data, waitter *sync.WaitGroup) {
	defer waitter.Done()
	res, err := http.Get(url)
	if err != nil {
		fetchedData.mu.Lock()
		fetchedData.d[url] = "error can't get the url"
		fetchedData.mu.Unlock()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fetchedData.mu.Lock()
		fetchedData.d[url] = "error can't get the url"
		fetchedData.mu.Unlock()
	}

	fetchedData.mu.Lock()
	fetchedData.d[url] = string(body)
	fetchedData.mu.Unlock()
}
