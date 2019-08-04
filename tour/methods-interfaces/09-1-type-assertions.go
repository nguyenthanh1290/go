package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

/*
A type assertion provides access to an interface value's underlying concrete value.

t := i.(T)
THIS STATEMENT ASSERTS THAT THE INTERFACE VALUE I HOLDS THE CONCRETE TYPE T AND ASSIGNS THE UNDERLYING T VALUE TO THE VARIABLE T.

IF I DOES NOT HOLD A T, THE STATEMENT WILL TRIGGER A PANIC.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.
*/
