package main

import (
	"fmt"
	"math/rand"
)

func cleanup() {
	fmt.Println("Cleaning up...")
	fmt.Println("Done.")
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}

// https://talks.golang.org/2012/concurrency.slide#38
// Receive on quit channel
// How do we know it's finished? Wait for it to tell us it's done: receive on the quit channel
