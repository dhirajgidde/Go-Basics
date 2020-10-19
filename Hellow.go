package main

// fmt stands for a format
import "fmt"

func main() {
	/*
		types of variable:
		1. int
		2. float64
		3. bool
		4.String
		%T to type of varaible
		%v for a value
		%q for a string
	*/
	fmt.Println("Enter the first Number:")
	var num int
	fmt.Scanln(&num)

	var a float64 = 4
	var b float64 = 30
	fmt.Println(b / a)
	if num == 10 {
		fmt.Println("The Num is greater than 10")
	} else if num == 20 {
		fmt.Println("The Num is 20")
	} else {
		fmt.Println("The num is not 10 nor 20")
	}
	fmt.Println("Hellow World ", num)

	switch num {
	case 10:
		fmt.Println("THe num value is 10")
		break
	case 20:
		fmt.Println("The num value is 20")

	case 30:
		fmt.Println("THe number is 30")

	default:
		fmt.Println("Wrong Choice")
	}

	for i := 0; i < num; i++ {
		fmt.Println(i)
	}

	/*
		infinite for loop
		for i:=0; ; i++   for i:=0; ; i++
		for { }
	*/

	/*

	 */
}
