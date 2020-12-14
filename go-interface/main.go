package main

import "fmt"

type Bot interface {
	welcome(val interface{})
}

func welcome(val interface{}) {
	fmt.Println(val)
}

func main() {
	val1 := "Welcome to Goa Singham"
	val2 := 105
	val3 := 12.5

	welcome(val1)
	welcome(val2)
	welcome(val3)
}
