package main

import (
	"fmt"
	"net/http"
	"time"
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

	// fmt.Println(<-c) // this is blocking call until unless we recieve any value in the channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	/*
		for i := 0; i < len(links); i++ {
			fmt.Println(<-c) // this is blocking call until unless we recieve any value in the channel
		}
	*/

	/*
		for {
			go checkLink(<-c, c)
		}
	*/

	for l := range c {
		// time.Sleep(time.Second) // this is not the correct place to sleep as per defination of Sleep - it will sleep the main routine instead of child routine
		// go checkLink(l, c)

		// functional literal to rescure this situation which is equivalent to ananomys function in javascript.

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
	/**

	We're saying wait for the channel to return some value after the channel has returned some value.
	Assign it to this variable l in this case being short for link then run the body of the for loop and
	inside the for loop we immediately spawn a new go routine calling check link passing in the link that
	we've just received in the channel and then passing the channel as the second argument.

	*/
}

/*
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
*/

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- "Might be down link"
		return
	}

	fmt.Println(link, " is up")
	c <- link
}
