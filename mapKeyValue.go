package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["4"] = 5
	m["3"] = 6
	m["2"] = 7

	values := make([]int, len(m))
	keys := make([]string, len(m))

	i := 0
	for k, v := range m {
		values[i] = v
		keys[i] = k
		i++
	}

	fmt.Println(keys, values)

}
