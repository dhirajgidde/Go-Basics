package main

import "fmt"

func demo(a, b int) int {
	fmt.Println("Welcome to demo function")
	val1 := a * b
	return val1
}

func main() {

	chk := 5
	fmt.Println(chk)
	// for i := 0; i < chk; i++ {
	// 	fmt.Println(i)
	// }

	// for {
	// 	fmt.Println("We are in Infinite loop")
	// }

	// for i := 0; ; {
	// 	fmt.Println("We are here")
	// }

	for chk < 10 {
		fmt.Println(chk)
		chk++
	}

	dd := demo(10, 20)
	fmt.Println(dd)

	//anonymous function
	add := func(a, b int) int {
		var c int
		c = a + b
		return c
	}

	fmt.Println(add(20, 20))
}
