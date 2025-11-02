package main

import (
	"fmt"
	"foundations/calculate"
)

func main() {
	result := calculate.Sum(1, 33)

	fmt.Println("The sum is: ", result)
}
