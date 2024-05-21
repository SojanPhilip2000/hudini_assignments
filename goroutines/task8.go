 // Objective:
// Learn how to use the select statement to handle multiple channels.
 
// Task:
 
// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.
//has context menu


package main

import (
	"fmt"

)

func sendmsg1(msg chan string){
		msg <- "hi hello"
}

func sendmsg2(msg chan string){
	msg <- "how are you"
}


func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)


	go sendmsg1(ch1)
	go sendmsg2(ch2)

	for i:=0; i<2; i++{
		select{
			case val:= <-ch1:
				fmt.Println(val)
			
			case val:= <-ch2:
				fmt.Println(val)
			
			}
		}
}