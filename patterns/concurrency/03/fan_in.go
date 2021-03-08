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

func fanIn(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-in1
		}
	}()
	go func() {
		for {
			c <- <-in2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Jack"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

// https://talks.golang.org/2012/concurrency.slide#27
// We can instead use a fan-in function to let whosoever is ready talk.