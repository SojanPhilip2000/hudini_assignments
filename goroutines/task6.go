// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:

// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.

package main

import (
	"time"
	"fmt"

)

func sendmsg(message []string,ch chan string){
	for _,msg:= range message{
		ch <- msg
		time.Sleep(time.Second)
	}

}


func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	message1:= []string{"hi", "how", "are","you"}
	message2:= []string{"my", "name", "is","sojan"}

	go sendmsg(message1, ch1)
	go sendmsg(message2, ch2)

	for {
		select{
		case msg:= <-ch1:
			fmt.Println("channel 1 is ", msg)
		
		case msg:= <-ch2:
			fmt.Println("channel 2 is ", msg)
		default:
			fmt.Println("no message is received")
			time.Sleep(time.Second)
		}
		}
}

