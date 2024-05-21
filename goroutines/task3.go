// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.

package main

import (
	"fmt"
)

func dispmsg1(msg chan string) {
	msg <- "hello world"

}

func dispmsg2(msg chan string) {
	msg <- "go lang"
}

func main() {
	msg := make(chan string)
	go dispmsg1(msg)
	go dispmsg2(msg)

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
