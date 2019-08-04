package main

import (
	"fmt"
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	sfdsds
	numbdderOfDays  // this constant is not exported
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(Sunday)
	fmt.Println(Wednesday)
	fmt.Println(Partyday)
	fmt.Println(numbdderOfDays)
}
