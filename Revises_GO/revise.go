package main

import (
	"fmt"
	"time"
)

func check(str string) (num int, str1 string) {
	str1 = str
	num = 10
	return
}

func add(ch chan int) {
	nsd := <-ch
	fmt.Println(100 + nsd)
}

func main() {
	fmt.Println("Welcome to Go Again")

	var comp1 complex64

	comp1 = 24 + 2i
	comp2 := 28 + 5i

	fmt.Println(comp2)
	fmt.Println(comp1)

	var num1 int
	fmt.Println("Enter the Value")
	fmt.Scanln(&num1)
	fmt.Println(num1)

	var str string
	str = "Welcome"

	var rn rune
	rn = '!'
	fmt.Println(rn)

	chcc, _ := check(str)
	fmt.Println(chcc)

	ch := make(chan int)
	go add(ch)
	ch <- 100
	time.Sleep(1 * 1000)

}
