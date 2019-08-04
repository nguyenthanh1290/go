package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	// even though v is a value and not a pointer, the method with the pointer receiver is called automatically.
	// That is, as a convenience, Go interprets the statement v.Scale(10) as (&v).Scale(10) since the Scale method has a pointer receiver.

	fmt.Println(v.Abs())
}

// You can declare methods with pointer receivers.
// Methods with pointer receivers can modify the value to which the receiver points
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
