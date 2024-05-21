 // Objective:
// Understand how to use channels for communication between goroutines.
 
// Task:
 
// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:
 
// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.

package main

import (
	"fmt"
)

func PrintMessage(num chan int){
	for i:= 1; i<=10; i++{
		num <- i // Sends the value of i to the channel num
	}
	close(num) // Closes the channel num when all values have been sent
}


func main(){
	num := make(chan int) // Creates an unbuffered channel of type int
	go PrintMessage(num) // Starts a goroutine to execute the PrintMessage function

	for n := range num { // Iterates over the values received from the channel num
		fmt.Println(n) // Prints the received value
	}
}

 
 
