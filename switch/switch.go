package main

import "fmt"

func add() {
	fmt.Println(2 + 3)
}

func sub() {
	fmt.Println(2 - 3)
}

func mult() {
	fmt.Println(2 * 3)
}

func main() {

	var num int

	for true {

		fmt.Println("Enter the number:")
		fmt.Scanln(&num)

		switch num {
		case 1:
			add()
			break

		case 2:
			sub()
			break
		case 3:
			mult()
			break

		case 4:
			return

		}

	}

}
