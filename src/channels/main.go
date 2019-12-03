package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to go Channels")

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	fmt.Println(links)

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		return
	}

	fmt.Println(link, " is up")
}
