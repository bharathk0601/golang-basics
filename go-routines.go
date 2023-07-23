package main

import "sync"

var wg sync.WaitGroup

func main() {
	urls := []string{"https://google.com", "https://github.com"}

	for _, v := range urls {
		wg.Add(1)
		go makeGetReq(v)

	}

	wg.Wait()
}

func makeReq(url string) {
	defer wg.Done()
	makeGetReq(url)
}
