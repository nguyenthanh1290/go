// Package triangle provides a function that determines triangle type.
package triangle

import "math"

// Kind defines the triangle type.
type Kind int

// Different kinds of triangle
// - NaT: not a triangle
// - Equ: equilateral
// - Iso: isosceles
// - Sca: scalene
const (
	NaT Kind = 1 << iota
	Equ
	Iso
	Sca
)

// KindFromSides takes the given sides and returns the type of a triangle they can form.
func KindFromSides(a, b, c float64) Kind {
	// for _, v := range []float64{a, b, c} {
	// 	if v <= 0 || math.IsNaN(v) || math.IsInf(v, 1) {
	// 		return NaT
	// 	}
	// }
	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a+b+c) || math.IsInf(a+b+c, 1) {
		return NaT
	}

	if (a+b < c) || (a+c < b) || (b+c < a) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
