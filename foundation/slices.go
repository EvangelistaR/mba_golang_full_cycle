package main

import "fmt"

func main() {
	ids := []int{2, 4, 6, 8, 10}

	fmt.Printf("length=%d capability=%d %v\n", len(ids), cap(ids), ids)

	fmt.Printf("length=%d capability=%d %v\n", len(ids[:1]), cap(ids[:1]), ids[:1])

	fmt.Printf("length=%d capability=%d %v\n", len(ids[3:]), cap(ids[3:]), ids[3:]) // degrease slice's capability

	ids = append(ids, 12) // doubling the slice's capability

	for index := range ids {
		fmt.Printf("ID: %d\n", ids[index])
	}
}
