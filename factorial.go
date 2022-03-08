package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)

	fact_rec := factorial_rec(5)
	fmt.Println(fact_rec)

}

func factorial(n int) int {
	var total int = 1
	for ; n > 1; n-- {
		total *= n

	}
	return total
}

func factorial_rec(n int) int {

	if n == 1 {
		return 1
	}

	return n * factorial_rec(n-1)

}
