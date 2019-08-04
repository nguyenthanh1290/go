// Package raindrops provides a function that convert raindrops to its appropriate sound.
package raindrops

import "fmt"

// Convert returns a sound based on the number of raindrops.
func Convert(drops int) string {
	s := ""
	if drops%3 == 0 {
		s += "Pling"
	}
	if drops%5 == 0 {
		s += "Plang"
	}
	if drops%7 == 0 {
		s += "Plong"
	}

	if s == "" {
		s = fmt.Sprintf("%v", drops)
	}

	return s
}
