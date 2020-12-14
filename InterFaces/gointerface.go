package main

import "fmt"

/*
	In Go lang Interface is used as a generic variable
*/
type tank interface {
	area(a, b int) (val int)
}

func area(a, b int) (val int) {
	val = a * b
	return
}

//This is called as a polymorphism
func myfun(a interface{}) {

	// Extracting the value of a
	val := a
	fmt.Println("Value: ", val)
}

func main() {
	//	var t tank

	fmt.Println(area(10, 20))

	//	fmt.Println(t.area(10, 20))

	var val interface {
	} = "GeeksforGeeks"

	myfun(val)
	myfun(123)

}
