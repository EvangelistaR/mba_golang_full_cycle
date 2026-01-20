package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := sum(48, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

// set a function to receive 2 integers, sum up them, return a new one and throw an error
func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("Cannot sum up the values")
	}
	return a + b, nil
}
