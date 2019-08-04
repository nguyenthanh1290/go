package main

import (
	"fmt"
	"strings"
)

const (
	nothing    = `NOTHING`
	add        = `ADD %s`
	change     = `CHANGE %s %s`
	move       = `MOVE %s`
	impossible = `IMPOSSIBLE`
)

func equal(S, T string) string {
	if S == T {
		return nothing
	}
	return ""
}

func addable(S, T string) string {
	if len(S) == len(T)-1 && S == T[:len(T)-1] {
		return fmt.Sprintf(add, T[len(T)-1:])
	}

	return ""
}

func changable(S, T string) string {
	n := len(S)
	m := len(T)

	if n == m {
		diffIndexes := []int{}
		i := 0
		for i < n {
			if S[i] != T[i] {
				diffIndexes = append(diffIndexes, i)
				if len(diffIndexes) > 1 {
					return ""
				}
			}
			i++
		}
		return fmt.Sprintf(change, string(S[diffIndexes[0]]), string(T[diffIndexes[0]]))
	}

	return ""
}

func moveable(S, T string) string {
	n := len(S)
	m := len(T)

	if n == m {
		for i := 0; i < n; i++ {
			if S[i] != T[i] {
				s := S[i+1:]
				t := strings.Replace(T[i:], string(S[i]), "", 1)
				if s == t {
					return fmt.Sprintf(move, string(S[i]))
				}
				return ""
			}
		}
	}

	return ""
}

func obtain(S, T string) string {
	S = strings.ToLower(S)
	T = strings.ToLower(T)

	result := equal(S, T)
	if result != "" {
		return result
	}

	result = addable(S, T)
	if result != "" {
		return result
	}

	result = changable(S, T)
	if result != "" {
		return result
	}

	result = moveable(S, T)
	if result != "" {
		return result
	}

	return impossible
}

func main() {
	fmt.Println(obtain("", ""))             // NOTHING
	fmt.Println(obtain("a", "a"))           // NOTHING
	fmt.Println(obtain("a", "A"))           // NOTHING
	fmt.Println(obtain("abcd", "AbCD"))     // NOTHING
	fmt.Println(obtain("0", "odd"))         // IMPOSSIBLE
	fmt.Println(obtain("nice", "nicer"))    // ADD r
	fmt.Println(obtain("xx", "xXx"))        // ADD x
	fmt.Println(obtain("test", "tent"))     // CHANGE s n
	fmt.Println(obtain("text", "tent"))     // CHANGE x n
	fmt.Println(obtain("text", "texx"))     // CHANGE t x
	fmt.Println(obtain("test", "tenk"))     // IMPOSSIBLE
	fmt.Println(obtain("beans", "banes"))   // MOVE e
	fmt.Println(obtain("beanse", "banese")) // MOVE e
}
