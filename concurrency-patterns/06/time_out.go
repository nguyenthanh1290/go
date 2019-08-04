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

func main() {
	c := boring("Long")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(500 * time.Millisecond):// timeout for each message, that's why it takes too long to time out.
			fmt.Println("You're too slow.")
			return
		}
	}
}

// https://talks.golang.org/2012/concurrency.slide#35
