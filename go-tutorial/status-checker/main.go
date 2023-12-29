package main

import (
	"fmt"
	"net/http"
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
	
	// infinite loop
	for  {
		go checkLink(<-c, c)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}
	fmt.Println(url, "is up")
	c <- url
}