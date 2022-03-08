package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string
	Lname string
	Age   int
}

func main() {

	p1 := person{
		Name:  "James",
		Lname: "Bond",
		Age:   32}

	p2 := person{
		Name:  "Miss",
		Lname: "Moneypenny",
		Age:   27}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
