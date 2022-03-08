package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByAge) Sort_by_tarlan() {

	for i, j := 0, 1; i < len(a); i++ {

		if !a.Less(i, j) {
			for x := 0; x < j; x++ {

				if a.Less(j, x) {
					a := append(a[:x], a[j], a[x:])

					a.Swap(x, j)

				}

			}

		}

	}
}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
