package main

import "fmt"

type person struct {
	name   string
	age    int
	salary float64
}

func (p person) show() {
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.salary)
}

func main() {

	persons := person{
		name:   "Dhiraj Gidde",
		age:    20,
		salary: 450.100,
	}

	student := struct {
		name  string
		age   int
		marks float64
	}{
		name:  "Dhiraj Dada",
		age:   23,
		marks: 56.25,
	}

	persons.show()

	fmt.Println(student)

	anonymous := func(val1 int) (val2 int) {
		val2 = 10 + val1
		return
	}(20)

	fmt.Println(anonymous)

}
