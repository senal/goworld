package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.amazon.com",
		"https://www.golang.com",
	}
	c := make(chan string, len(links))

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go (func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		})(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is up!")
	c <- link
}
