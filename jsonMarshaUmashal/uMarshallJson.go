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
	jsonBlub := []byte(`[{"Name":"James","Lname":"Bond","Age":32},{"Name":"Miss","Lname":"Moneypenny","Age":27}]`)
	people := []person{}

	err := json.Unmarshal(jsonBlub, &people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("value: %v, type: %T\n", people, people)
}
