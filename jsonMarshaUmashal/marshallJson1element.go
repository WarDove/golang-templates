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
	tarlan := person{
		Name:  "tarlan",
		Lname: "huseynov",
		Age:   30,
	}

	myjson, err := json.Marshal(tarlan)

	if err != nil {
		fmt.Println(err)
	}

	println(string(myjson))

}
