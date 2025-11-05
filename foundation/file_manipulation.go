package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const FILENAME = "file-example.txt"
	f, err := os.Create(FILENAME)
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("That's my first note!\n"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created file with success!\n Size: %d bytes", size)

	file, err := os.ReadFile(FILENAME)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	file2, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove(FILENAME)

	if err != nil {
		panic(err)
	}
}
