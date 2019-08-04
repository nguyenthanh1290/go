package main

import (
	"sync"
	"fmt"
)

type FreqMap map[rune]int

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

func main() {
	var f =	ConcurrentFrequency([]string{euro, dutch, us})

	fmt.Println(f)
}

var (
	euro = `Freude schöner Götterfunken
Tochter aus Elysium,
Wir betreten feuertrunken,
Himmlische, dein Heiligtum!
Deine Zauber binden wieder
Was die Mode streng geteilt;
Alle Menschen werden Brüder,
Wo dein sanfter Flügel weilt.`

	dutch = `Wilhelmus van Nassouwe
ben ik, van Duitsen bloed,
den vaderland getrouwe
blijf ik tot in den dood.
Een Prinse van Oranje
ben ik, vrij, onverveerd,
den Koning van Hispanje
heb ik altijd geëerd.`

	us = `O say can you see by the dawn's early light,
What so proudly we hailed at the twilight's last gleaming,
Whose broad stripes and bright stars through the perilous fight,
O'er the ramparts we watched, were so gallantly streaming?
And the rockets' red glare, the bombs bursting in air,
Gave proof through the night that our flag was still there;
O say does that star-spangled banner yet wave,
O'er the land of the free and the home of the brave?`
)