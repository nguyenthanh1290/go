// Package secret provides functions that decodes the code to the appropriate sequence of events for a secret handshake.
package secret

const (
	wink = 1 << iota
	doubleBlink
	closeEyes
	jump
	reverse
)

func reverseArray(input []string) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

// Handshake returns a decoded handshake.
func Handshake(code uint) (output []string) {
	if code&wink != 0 {
		output = append(output, "wink")
	}
	if code&doubleBlink != 0 {
		output = append(output, "double blink")
	}
	if code&closeEyes != 0 {
		output = append(output, "close your eyes")
	}
	if code&jump != 0 {
		output = append(output, "jump")
	}
	if code&reverse != 0 {
		reverseArray(output)
	}
	return
}
