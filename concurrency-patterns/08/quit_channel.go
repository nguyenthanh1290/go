package main

import (
	"fmt"
	"math/rand"
)

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				return
			}
		}
	}()
	return c
}

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(40); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

// https://talks.golang.org/2012/concurrency.slide#37
