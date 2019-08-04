package main

import "fmt"

func main() {
	// int array
	var arr [3]int
	arr[0] = 1
	arr[0] = 2
	arr[0] = 3

	// alternative syntax
	arr2 := [4]int{3, 2, 5, 1}
	fmt.Println(arr2)

	//string array
	var strArray [3]string
	strArray[0] = "hello"
	strArray[1] = "hi"
	strArray[2] = "hey"
	fmt.Println(strArray[2])

	arrp := []int{1, 3, 6, 2, 6}
	arrp = append(arrp, 23)
	fmt.Println(arrp)
}