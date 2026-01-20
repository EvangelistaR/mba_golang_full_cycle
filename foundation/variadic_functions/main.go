package main

import "fmt"

func main() {
	fmt.Println(sum(1,3,45, 6,34))
}

func sum(numbers ...int) int {
	amount := 0
	for _, number := range numbers {
		amount += number
	}

	return amount
}