package main

import "fmt"

func fib(num int) {
	a := 0
	b := 1
	for i := 0; i <= num; i++ {
		temp := a + b
		a = b
		b = temp
		fmt.Print(b, " ")
	}
}

func sum_of_n(arr []int, c chan int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	c <- sum
}

func main() {

	fib(10)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)

	go sum_of_n(arr, ch)
	x := <-ch

	fmt.Println("\nThe sum of all n numbers")
	fmt.Println(x)

}
