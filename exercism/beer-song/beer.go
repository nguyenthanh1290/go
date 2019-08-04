package beer

import (
	"errors"
	"fmt"
	"strings"
)

const (
	upperBound int = 99
	lowerBound int = 0
)

const (
	v0 = `No more bottles of beer on the wall, no more bottles of beer.\n
	Go to the store and buy some more, 99 bottles of beer on the wall.\n`

	v1 = `1 bottle of beer on the wall, 1 bottle of beer.\n
	Take it down and pass it around, no more bottles of beer on the wall.\n`

	v2 = `%v bottles of beer on the wall, %v bottles of beer.\n
	\Take one down and pass it around, 1 bottle of beer on the wall.\n`

	vN = `%v bottles of beer on the wall, %v bottles of beer.\n
	Take one down and pass it around, %v bottles of beer on the wall.\n`
)

// Verse returns the verse based on the input bottle.
func Verse(bottle int) (string, error) {
	if bottle < lowerBound || bottle > upperBound {
		return "", errors.New("invalid number of bottles")
	}

	if bottle == 0 {
		return v0, nil
	}
	if bottle == 1 {
		return v1, nil
	}
	if bottle == 2 {
		return fmt.Sprintf(v2, bottle, bottle), nil
	}
	return fmt.Sprintf(vN, bottle, bottle, bottle-1), nil
}

// Verses returns a range from [upper] to [lower] of verses.
func Verses(upper, lower int) (string, error) {
	if lower < lowerBound || upper > upperBound || lower > upper {
		return "", errors.New("invalid bound(s)")
	}

	verses := make([]string, upper-lower+1)
	for i := upper; i >= lower; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses = append(verses, v, "\n")
	}

	return strings.Join(verses, ""), nil
}

// Song returns the entire song.
func Song() string {
	s, _ := Verses(99, 0)
	return s
}
