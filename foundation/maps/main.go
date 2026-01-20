package main

import "fmt"

func main() {
	// setting payments as map with key and value
	payments := map[string]int{"Wesley": 1_000, "John": 2_000, "Mariah": 5_000}

	// accessing a specific key and display the value
	fmt.Println(payments["Mariah"])

	// remove an item from payments map
	delete(payments, "John")
	// set a new item to payments map
	payments["Johny"] = 2_300
	fmt.Println(payments["Johny"])

	// listing the name and amount
	for key, value := range payments {
		fmt.Printf("Name: %s\nAmount: %d\n", key, value)
	}

	// another way to set a map
	salaries := make(map[string]int) // or just salaries := map[string]int{}
	salaries["Jane"] = 4_300

	for key, value := range salaries {
		fmt.Printf("Name: %s\nAmount: %d\n", key, value)
	}

	// listing only the values
	for _, value := range payments {
		fmt.Printf("$ %d\n", value)
	}
}
