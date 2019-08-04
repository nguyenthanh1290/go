package main

import (
	"fmt"

	"github.com/nguyenthanh1290/codility/sorting/bubble"
	"github.com/nguyenthanh1290/codility/sorting/bucket"
	"github.com/nguyenthanh1290/codility/sorting/insertion"
	"github.com/nguyenthanh1290/codility/sorting/merge"
	"github.com/nguyenthanh1290/codility/sorting/selection"
)

func main() {
	fmt.Print("Un-sorted: ")
	a := []int{3, 2, 4, 2, 3, 5}
	fmt.Println(a)
	fmt.Println(bucket.SortInts(a))
	fmt.Println(bubble.SortInts(a))
	fmt.Println(selection.SortInts(a))
	fmt.Println(insertion.SortInts(a))
	fmt.Println(merge.SortInts(a))
	fmt.Println()

	fmt.Print("Negative un-sorted: ")
	a = []int{3, 2, -4, 2, 0, 5, 0, -1, 200}
	fmt.Println(a)
	fmt.Println(bucket.SortInts(a))
	fmt.Println(bubble.SortInts(a))
	fmt.Println(selection.SortInts(a))
	fmt.Println(insertion.SortInts(a))
	fmt.Println(merge.SortInts(a))
	fmt.Println()

	fmt.Print("Sorted: ")
	a = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(bucket.SortInts(a))
	fmt.Println(bubble.SortInts(a))
	fmt.Println(selection.SortInts(a))
	fmt.Println(insertion.SortInts(a))
	fmt.Println(merge.SortInts(a))
	fmt.Println()

	fmt.Print("Reverse sorted: ")
	a = []int{5, 4, 3, 2, 1, 0}
	fmt.Println(a)
	fmt.Println(bucket.SortInts(a))
	fmt.Println(bubble.SortInts(a))
	fmt.Println(selection.SortInts(a))
	fmt.Println(insertion.SortInts(a))
	fmt.Println(merge.SortInts(a))
	fmt.Println()
}
