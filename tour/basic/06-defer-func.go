package main

import (
	"fmt"
)

func main() {
	i := c()
	fmt.Println(i)
}

func c() (i int) {
    defer func() { i++ }()
    return 1
}

// 3. Deferred functions may read and assign to the returning function's named return values.