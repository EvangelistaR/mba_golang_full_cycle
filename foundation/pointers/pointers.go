package main

func main() {
	// memory -> address -> value
	a := 10
	var pointer *int = &a

	println(pointer)

	*pointer = 20
	println(&a)

	b := &a

	println(*b)
	*b = 30
	println(a)
}
