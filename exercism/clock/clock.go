// Package clock implements a clock that handles times without dates.
package clock

import "fmt"

// Clock defines the clock type with hour and minute.
type Clock struct {
	hour, minute int
}

const minOfDay = 24 * 60

// New returns a new clock.
func New(h, m int) Clock {
	totalMin := (minOfDay + (h*60+m)%minOfDay) % minOfDay
	h = totalMin / 60
	m = totalMin % 60

	return Clock{h, m}
}

// Add returns a new clock after adding minutes to it.
func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

// Subtract returns a new clock after subtracting minutes to it.
func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

// String returns a string that represents the clock in format of [hh]:[mm].
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
