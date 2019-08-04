package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency is the concurrent version of Frequency.
func ConcurrentFrequency(strs []string) FreqMap {
	var counter = struct {
		sync.Mutex
		freq FreqMap
	}{freq: map[rune]int{}}

	var wg sync.WaitGroup
	wg.Add(len(strs))

	for _, s := range strs {
		go func(s string) {
			for _, r := range s {
				counter.Lock()
				counter.freq[r]++
				counter.Unlock()
			}
			wg.Done()
		}(s)
	}

	wg.Wait()

	return counter.freq
}
