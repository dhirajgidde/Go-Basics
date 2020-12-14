package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(link + " is up")
		ch <- "Yeah it's working"

	}
	//ch <- "Yep, its working"
}

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://www.twitter.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
}
