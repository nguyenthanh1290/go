// A concurrent prime sieve

package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		fmt.Print("i=", i)
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println("prime=", prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

/*
i=2 i=3
prime= 2
prime= 3
i=4 i=5 i=6 i=7 i=8 i=9 i=10
prime= 5
prime= 7
i=11 i=12 i=13 i=14 i=15 i=16 i=17 i=18 i=19 i=20 i=21 i=22
prime= 11
prime= 13
prime= 17
i=23 i=24 i=25 i=26 i=27 i=28 i=29 i=30 i=31 i=32 i=33 i=34 i=35 i=36 i=37 i=38 i=39 i=40 i=41 i=42 i=43 i=44 i=45 i=46
prime= 19
prime= 23
prime= 29
i=47
*/