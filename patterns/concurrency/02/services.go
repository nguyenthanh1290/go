package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boring returns receive-only channel of strings.
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func main() {
	joe := boring("Joe")
	jack := boring("Jack")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-jack)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

// https://talks.golang.org/2012/concurrency.slide#26
// Channels as a handle on a service
// These programs make Joe and Ann count in lockstep.