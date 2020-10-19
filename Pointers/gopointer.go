package main

import "fmt"

func prtt(ptr *int) {
	*ptr = 100
}

func updateSlice(arr []int) {
	arr[4] = 10
}

func main() {

	var num = 10

	var s *int

	s = &num

	var ss *int = &num

	var s1 **int = &s

	fmt.Println("Num:", *(&num))
	fmt.Println("Pointer s", *s)
	fmt.Println("Address of num:", *(*(&s)))
	fmt.Println("Address of s: ", *(*(*(&s1))))
	prtt(s)
	fmt.Println("Value of Num:", num)
	fmt.Println(&s == &ss)

	slc := []int{1, 2, 3, 4, 5, 6}

	updateSlice(slc)

	fmt.Println(slc)
}
