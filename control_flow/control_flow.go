package main

import "fmt"

func main() {

	fmt.Println("First")
	defer fmt.Println("Second")
	fmt.Println("Third")
	panic("Something is wrong here")

}
