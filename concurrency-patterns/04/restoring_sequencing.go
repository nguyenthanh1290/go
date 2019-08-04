package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Message defines a message type that contains a channel for the reply.
type Message struct {
	str  string
	wait chan bool
}

// https://talks.golang.org/2012/concurrency.slide#29
// Send a channel on a channel, making goroutine wait its turn.
// Receive all messages, then enable them again by sending on a private channel.