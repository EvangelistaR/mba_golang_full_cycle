package main

import "fmt"

// type ID int

// var (
// 	identificator ID = 10
// )

func main() {
	type ID int
	var (
		identificator ID = 10
	)
	fmt.Printf("identificator type is: %T", identificator)
	fmt.Printf("\nidentificator value is: %v\n", identificator)
}
