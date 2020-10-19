package main

import "fmt"

type product struct {
	id      int
	name    string
	price   float64
	details sheeping
}

type sheeping struct {
	ids int
	qty float64
}

type check struct {
	id   int
	name string
}

func (p product) show() (val int) {
	fmt.Println("Id: ", p.id)
	fmt.Println("Name: ", p.name)
	fmt.Println("price: ", p.price)
	fmt.Println("Sheeping ID: ", p.details.ids)
	fmt.Println("QTY:", p.details.qty)
	val = 5
	return
}

func (p product) totalprice() (prices float64) {
	prices = p.price * p.details.qty
	return
}

func ano() {

	//Anonymous Class
	Element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Pikachu",
		branch:    "ECE",
		language:  "C++",
		Particles: 498,
	}

	fmt.Println(Element)
}

func main() {

	var chk [3]check

	for i := 0; i < 3; i++ {
		var id int
		var name string
		fmt.Println("Enter the id:")
		fmt.Scanln(&id)
		fmt.Println("Enter the name:")
		fmt.Scanln(&name)
		chk[i].id = id
		chk[i].name = name
	}

	fmt.Println(chk)

	prod := product{
		id:      1,
		name:    "Banana",
		price:   20.4,
		details: sheeping{2, 5},
	}

	fmt.Println(prod)
	fmt.Println(prod.show())
	ano()
	fmt.Println(prod.totalprice())
}
