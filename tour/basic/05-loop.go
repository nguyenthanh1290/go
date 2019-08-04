package main

import "fmt"

func main() {
	// loops
	fmt.Println("I'm looping...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum: ", sum)

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// forever
	// for {
	// }

	// loop + array
	arrl := []string{"a", "b", "c"}
	for index, value := range arrl {
		fmt.Println("index: ", index, " - value: ", value)
	}

	// loop + map
	mapl := make(map[string]string)
	mapl["a"] = "alpha"
	mapl["b"] = "beta"
	
	for key, value := range mapl {
		fmt.Println("key: ", key, " - value: ", value)
	}
}
