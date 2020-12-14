package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string `json:"name"`
	Sales int    `json:"sales"`
	Age   int    `json:"age"`
}

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

func main() {
	author := Author{Name: "Shivaji Sawant", Sales: 100, Age: 45}
	book := Book{Title: "Chhava", Author: author}
	//fmt.Println(book)
	//fmt.Printf("Book:%+v", book)

	//convert a struct into a json value
	//	bytearr, err := json.Marshal(book)
	bytearr, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytearr))
}
