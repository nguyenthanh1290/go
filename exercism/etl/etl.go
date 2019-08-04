package etl

import (
	"strings"
)

// Transform returns an ETL of input.
func Transform(in map[int][]string) (out map[string]int) {
	out = map[string]int{}

	for k, v := range in {
		for _, e := range v {
			out[strings.ToLower(e)] = k
		}
	}

	return
}
