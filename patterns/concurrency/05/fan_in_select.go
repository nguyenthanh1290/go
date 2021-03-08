package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-in1:
				c <- s
			case s := <-in2:
				c <- s
			}
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

// https://talks.golang.org/2012/concurrency.slide#32
// The select statement provides another way to handle multiple channels. 
// It's like a switch, but each case is a communication: 
// - All channels are evaluated. 
// - Selection blocks until one communication can proceed, which then does. 
// - If multiple can proceed, select chooses pseudo-randomly. 
// - A default clause, if present, executes immediately if no channel is ready.