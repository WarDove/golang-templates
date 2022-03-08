package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string `json:"Name"`
	Lname string
	Age   int
}

func main() {
	jsonBlub := []byte(`{"Name":"James","Lname":"Bond","Age":32}`)
	p := person{}

	err := json.Unmarshal(jsonBlub, &p)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("value: %v, type: %T\n", p, p)
}
