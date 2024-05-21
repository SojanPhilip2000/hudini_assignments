package main

import (
	"fmt"
	"time"
)

type waitgroup struct {
	count int
}

func (wg *waitgroup) Add(num int) {
	wg.count += num
}

func (wg *waitgroup) Done() {
	wg.count--
}

func (wg *waitgroup) Wait() {
	for {
		if wg.count == 0{
			return
		}
	}
}

func Looper(name string, wg *waitgroup) {
	defer wg.Done()
	fmt.Println(name)

}
func main() {
	wg := waitgroup{}
	wg.Add(3)
	go Looper("routine1", &wg)
	go Looper("routine2", &wg)
	go Looper("routine2", &wg)
	wg.Wait()
}
