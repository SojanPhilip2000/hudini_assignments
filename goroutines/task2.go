// Objective:
// Learn how to create and use goroutines.

// Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.
package main

import (
	"fmt"
	"sync"
	"time"
)

// // func execute(name string){
// // 	for i:=1; i<=5; i++{
// // 		time.Sleep(time.Second)
// // 		fmt.Println(name,i)
// // 	}
// // }

// // func main(){
// // 	go execute("routine1")
// // 	go execute("routine2")
// // 	go execute("routine2")
// // }

// func execute(name string){
// 	for i:=1; i<=5; i++{
// 		time.Sleep(time.Second)
// 		fmt.Println(name,i)
// 	}
// }

// func main(){
// 	go execute("routine1")
// 	go execute("routine2")
// 	go execute("routine2")
// 	time.Sleep(time.Second)
// }

func execute(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(name, i)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go execute("routine1", wg)
	go execute("routine2", wg)
	go execute("routine2", wg)
	wg.Wait()
}


