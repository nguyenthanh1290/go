package main

import "fmt"

// A var statement can be at package or function level.
var c, python, java bool

func main() {
	fmt.Println(c, python, java)

	// A var declaration can include initializers, one per variable.
	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

	fmt.Println("integer:")
	var a int = 1
	var b, c int
	b = 2
	c = 3
	fmt.Println(a)
	fmt.Println(b + c)

	fmt.Println("floating point number:")
	var d float64 = 9.99
	fmt.Println(d)

	var str string = "This is a string"
	fmt.Println(str)

	// Shorthand Declaration
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	// foo := "bar" is equivalent to var foo string = "bar"
	e := 9
	f := "golang"
	g := 923.2

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Outside a function, every statement begins with a keyword (var, func, and so on)
	// and so the := construct is not available.
}
