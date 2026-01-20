package main

import "fmt"

type NewNumber int

type Number interface {
	~int | ~float64 // with that, NewNumber is being allowed as a type
}

func Sum[T Number](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {

	salaries1 := map[string]int{"Jonn": 1_000, "Mary": 5_300}
	salaries2 := map[string]float64{"John": 1_899.89, "Mary": 5_387.99}
	salaries3 := map[string]NewNumber{"John": 1_899, "Mary": 5_000}

	fmt.Println(Sum(salaries1))
	fmt.Printf("%.2f", Sum(salaries2))
	fmt.Println(Sum(salaries3))

	println(Compare(10, 10))
}
