package house

import (
	"bytes"
	"strings"
)

var (
	subjects = [...]string{"This", "that", "that", "that", "that", "that", "that", "that", "that", "that", "that", "that"}
	verbs    = [...]string{
		"is",
		"belonged to",
		"kept",
		"woke",
		"married",
		"kissed",
		"milked",
		"tossed",
		"worried",
		"killed",
		"ate",
		"lay in",
	}
	objects = [...]string{
		"the house that Jack built",
		"the malt",
		"the rat",
		"the cat",
		"the dog",
		"the cow with the crumpled horn",
		"the maiden all forlorn",
		"the man all tattered and torn",
		"the priest all shaven and shorn",
		"the rooster that crowed in the morn",
		"the farmer sowing his corn",
		"the horse and the hound and the horn",
	}
)

// Verse returns a part of the song.
func Verse(depth int) string {
	var buf bytes.Buffer

	// First line
	s := subjects[0]
	v := verbs[0]
	o := objects[depth-1]

	line := strings.Join([]string{s, v, o}, " ")
	buf.WriteString(line)

	// Second line and more
	for i := 1; i < depth; i++ {
		buf.WriteString("\n")

		s = subjects[i]
		v = verbs[len(verbs)-depth+i]
		o = objects[depth-1-i]

		line = strings.Join([]string{s, v, o}, " ")
		buf.WriteString(line)
	}

	buf.WriteString(".")

	return buf.String()
}

// Song returns the entire song.
func Song() string {
	var buf bytes.Buffer
	for i := 1; i <= len(verbs); i++ {
		buf.WriteString(Verse(i))
		buf.WriteString("\n\n")
	}
	return strings.TrimSuffix(buf.String(), "\n\n")
}
