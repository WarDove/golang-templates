package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByFirst []Person

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i].First[0] < a[j].First[0] }

func main() {
	p1 := Person{"Tarlan", 32}
	p2 := Person{"Bilal", 27}
	p3 := Person{"Celal", 64}
	p4 := Person{"Azer", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByFirst(people))
	fmt.Println(people)

}
