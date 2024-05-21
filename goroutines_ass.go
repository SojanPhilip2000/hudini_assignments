// Objective:
// Learn how to send and receive values using channels.
 
// Task:
 
// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// The main function should receive the message from the channel and print it.
// Hints:
 
// Use an unbuffered channel for simplicity.

package main
import (
    "fmt"
)
func display(message chan string){
	message <- "hello world"

}
func main(){

	message:= make(chan string)
	go display(message)
	fmt.Println(<-message)

}
 

 