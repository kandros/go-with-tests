package main

import "fmt"

type WebsiteChecker func(string) bool
type result struct {
	url     string
	isAlive bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.isAlive
	}

	return results
}

func main() {
	fmt.Println("hello")
}
