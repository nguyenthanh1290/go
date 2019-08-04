// Package twofer implements "One for you and one for me." logic.
package twofer

import "fmt"

// ShareWith returns the string "One for {name}, one for me.". If name is empty, "you" is used.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
