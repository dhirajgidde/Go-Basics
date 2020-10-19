package main

import "fmt"

func arrAsReff(*arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(arr[i])
	}
}

func main() {

	var arr [5]int
	var num int
	for i := 0; i < 5; i++ {
		fmt.Scanln(&num)
		arr[i] = num
	}

	fmt.Println("The values inside the arrays is:")
	arrAsReff(arr, 5)

}
