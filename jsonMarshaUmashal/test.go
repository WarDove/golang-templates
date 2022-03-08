package main

import (
	"fmt"
)

type person struct {
	Name  string
	Lname string
	Age   int
}

func main() {

	tarlan := person{
		Name:  "Tarlan",
		Lname: "Huseynov",
		Age:   30,
	}

	tarlan.typem()

	tarlan.pointerm()

}

func (p *person) pointerm() {
	fmt.Println("receiver is a pointer")
	p.typem()
}

func (p person) typem() {
	fmt.Println("receiver is a type")
}
