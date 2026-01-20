package main

import "fmt"

func main() {
	amount := func() int {
		return sum(1,3,45, 6,34) * 2
	}()

	fmt.Println(amount)
}

func sum(numbers ...int) int {
	amount := 0
	for _, number := range numbers {
		amount += number
	}

	return amount
}