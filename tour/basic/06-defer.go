package main

import "fmt"

func main() {
	i := 0

	i++
	defer fmt.Println("defer i = ", i)

	i++
	fmt.Println("i = ", i)
}

// i= 2
// defer i =  1

// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately,
// but the function call is not executed until the surrounding function returns.

// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.