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

	c := make(chan string)

	for _, link := range links {
		// checkLink(link)
		go checkLink(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- "Might be down link"
		return
	}

	fmt.Println(link, " is up")
	c <- "yup link is up!!!"
}
