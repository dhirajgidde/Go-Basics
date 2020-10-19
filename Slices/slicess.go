package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	myslice := arr[1:4]

	myslice = arr[0:4]

	fmt.Println(arr)
	fmt.Println(len(myslice))
}
