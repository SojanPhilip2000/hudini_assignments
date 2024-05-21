// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// Task:

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.
package main

import (
	"fmt"
	"sync"
)

func execute(message string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(message)
}

func main(){
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go execute("hi", wg)
	go execute("hello", wg)
	go execute("how are you", wg)
    wg.Wait()

} 


