package main

import "fmt"

func main() {
	state := make(map[string]int)
	state = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}
	state["AF"] = 10
	fmt.Println(state["A"])
	fmt.Println(state)
	delete(state, "AF")
	fmt.Println(state)

	val, ok := state["ASD"] //if key is not present then value is return false
	fmt.Println(val, ok)

	// ' _ '  is used to blank identifier

	fmt.Println(len(state))
}
