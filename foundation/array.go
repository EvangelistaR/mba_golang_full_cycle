package main

import "fmt"

func main() {
	var ids [3]int

	ids[0] = 13
	ids[1] = 25
	ids[2] = 30

	for index, value := range ids {
		fmt.Printf("The index %d has value: %d\n", index, value)
	}
}
