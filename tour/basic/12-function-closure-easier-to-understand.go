package main

import (
	"fmt"
)

func main() {
	f := inc()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("New...")

	g := inc()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func inc() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

/*
1
2
3
4
5
New...
1
2
3
4
5
*/
