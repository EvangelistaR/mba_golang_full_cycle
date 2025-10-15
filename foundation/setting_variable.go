package main

// global scope
const a = "Hello, world!" // const cannot have the value changed

var (
	b bool    = true           // by default returns false
	c int     = 10             // by default returns 0
	d string  = "global scope" // by default returns a empty string
	e float64 = 10.33          // by default returns +0.000000e+000
)

func main() {
	// local scope
	var f = "Local instance"
	// shorthand
	shorthand_variable := "shorthand way"
	println(c)
	println(f)
	println(shorthand_variable)
}
