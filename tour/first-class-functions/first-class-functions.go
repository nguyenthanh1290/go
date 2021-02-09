package main

import "fmt"

func main() {
	fmt.Println("First class functions")
	s1 := Student{
		Firstname: "Steven",
		Lastname:  "Swachb",
		Age:       33,
		Grade:     "12",
	}
	s2 := Student{
		Firstname: "Steven",
		Lastname:  "Tran",
		Age:       30,
		Grade:     "10",
	}
	s3 := Student{
		Firstname: "Jonh",
		Lastname:  "Awd",
		Age:       30,
		Grade:     "10",
	}

	s := []Student{s1, s2, s3}
	f := filter(s, func(s Student) bool {
		if s.Age == 30 {
			return true
		}
		return false
	})

	fmt.Println("Filter")
	fmt.Println(f)

	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})

	fmt.Println("Mapping")
	fmt.Println(r)
}

// Student stores a student data
type Student struct {
	Firstname string
	Lastname  string
	Age       int
	Grade     string
}

func filter(s []Student, f func(s Student) bool) []Student {
	var r []Student

	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
