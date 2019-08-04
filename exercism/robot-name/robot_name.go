package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const max = 26 * 26 * 10 * 10 * 10

var counter = 0
var nameDB = map[string]bool{}

func makeName() string {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	c1 := rune('A' + rnd.Int()%26)
	c2 := rune('A' + rnd.Int()%26)
	n := rnd.Int() % 1000

	return fmt.Sprintf("%c%c%03d", c1, c2, n)
}

// Robot has a name.
type Robot struct {
	name string
}

// Name returns robot name.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if err := r.Reset(); err != nil {
			return "", err
		}
	}
	return r.name, nil
}

// Reset resets robot name.
func (r *Robot) Reset() error {
	if counter > max {
		return errors.New("namespace is exhausted")
	}

	name := makeName()
	for nameDB[name] {
		name = makeName()
	}

	nameDB[name] = true
	r.name = name
	counter++

	return nil
}
