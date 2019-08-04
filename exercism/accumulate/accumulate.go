package accumulate

// Accumulate returns a new collection containing the result of applying that operation to each element of the input collection.
func Accumulate(in []string, fn func(string) string) []string {
	out := make([]string, len(in))
	for i, e := range in {
		out[i] = fn(e)
	}
	return out
}
