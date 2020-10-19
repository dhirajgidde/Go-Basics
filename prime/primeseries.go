package main

import "fmt"

func primenumber(n int) {
	var arr [150]int
	index := 0
	i := 2
	for {
		j := 2
		for j < i {
			if i%j == 0 {
				break
			}
			j++
		}
		if i == j && i%1 == 0 {
			fmt.Print(i, " ")
			arr[index] = i
			index++
		}
		if index == n {
			break
		}
		i++

	}

	fmt.Println(arr[11/2])
}

func fibb(n int) {
	var arr [150]int
	index := 2
	temp := 0
	first := 0
	second := 1
	fmt.Print("\n")
	arr[0] = first
	arr[0] = second

	for i := 3; i <= n; i++ {
		temp = first + second
		first = second
		second = temp
		fmt.Print(second, " ")
		arr[index] = second
		index++
	}
	fmt.Println(arr[12/2])
}

func main() {
	primenumber(10)
	fibb(10)
}
