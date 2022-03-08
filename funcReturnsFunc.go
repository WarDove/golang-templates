package main

import "fmt"

func main() {

	oddtest := evenOrOddTrue("odd")
	fmt.Printf("5 is odd statement is %v\n", oddtest(5))
	fmt.Printf("6 is not odd statement is %v\n", oddtest(6))

	eventest := evenOrOddTrue("even")
	fmt.Printf("5 is not even statement is %v\n", eventest(5))
	fmt.Printf("6 is even statement is %v\n", eventest(6))
}

func evenOrOddTrue(numType string) func(int) bool {

	switch numType {

	case "odd":
		return func(i int) bool {

			if i%2 == 0 {
				return false
			} else {
				return true
			}
		}
	case "even":
		return func(i int) bool {

			if i%2 == 0 {
				return true
			} else {
				return false

			}
		}
	default:
		return func(i int) bool {
			fmt.Println("Wrong number type! False returned as default for any number!")
			return false

		}

	}

}
