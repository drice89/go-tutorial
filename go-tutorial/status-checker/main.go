package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const (
		facebookUrl = "http://facebook.com"
		googleUrl = "http://google.com"
		stackoverflowUrl = "http://stackoverflow.com"
		golangUrl = "http:/golang.org"
		amazonUrl = "http://amazon.com"
	)

	links := []string{
		facebookUrl, 
		googleUrl, 
		stackoverflowUrl, 
		golangUrl, 
		amazonUrl,
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}
	
	// range with a channel
	// wait for the channel to return some value, assign it to l
	// then run the block
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(url string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}
	fmt.Println(url, "is up")
	c <- url
}