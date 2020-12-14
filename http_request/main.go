package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(resp)
	}
}
