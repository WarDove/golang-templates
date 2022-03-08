package main

import "fmt"

func main() {

	s := map[int]int{
		1: 1, 2: 2, 3: 3,
	}

	if v, ok := s[1]; ok {
		fmt.Println(v, ok)
	}

}
