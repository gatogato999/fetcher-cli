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

	var waitter sync.WaitGroup
	var mu sync.Mutex
	for _, url := range urls {
		waitter.Add(1)
		go FetchData(url, &mu, &waitter)
	}

	waitter.Wait()
}

func FetchData(url string, mu *sync.Mutex, waitter *sync.WaitGroup) {
	defer waitter.Done()
	res, err := http.Get(url)
	if err != nil {
		mu.Lock()
		fmt.Print(err)
		mu.Unlock()
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		mu.Lock()
		fmt.Print(err)
		mu.Unlock()
		return
	}
	mu.Lock()
	fmt.Print(string(body))
	mu.Unlock()
}
