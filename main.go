package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	urls := os.Args[1:]
	// urls := []string{
	// 	"http://localhost:10055/marks/login",
	// 	"http://localhost:10055/marks/add",
	// 	"http://localhost:10055/marks/show",
	// }

	fetchedData := make(map[string]string)
	var waitter sync.WaitGroup
	for _, url := range urls {
		waitter.Add(1)
		go FetchData(url, fetchedData, &waitter)
	}

	waitter.Wait()
	for _, i := range fetchedData {
		fmt.Print("\n-------------------\n", i, "\n-------------------\n")
	}
}

func FetchData(url string, fetchedData map[string]string, waitter *sync.WaitGroup) {
	defer waitter.Done()
	res, err := http.Get(url)
	if err != nil {
		fetchedData[url] = "error can't get the url"
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fetchedData[url] = "error can't get the url"
	}

	fetchedData[url] = string(body)
}
