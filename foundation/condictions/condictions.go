package main

func main() {
	number1 := 1
	number2 := 2

	if number1 > number2 {
		println(number1)
	} else {
		println(number2)
	}

	number1 = 320
	switch number1 {
	case 1:
		println("number 1")
	case 2:
		println("number 2")
	case 3:
		println("number 3")
	default:
		println("invalid")
	}

}
