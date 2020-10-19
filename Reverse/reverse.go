package main

import "fmt"

func rev(num int) (revs int) {
	digit := 0

	for num > 0 {
		digit = num % 10
		num = num / 10
		revs = revs*10 + digit
	}
	return
}

func main() {
	fmt.Println("Weclcome to our area", rev(12345))
}
