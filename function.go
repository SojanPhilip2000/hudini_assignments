package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func print() {
	fmt.Println(add(23, 45))
	fmt.Printf("%d\n", sub(63, 43))
}

func swap(x, y string) (string, string) {
	return y, x
}

func printswap() {
	var a, b = swap("hai", "hello")
	fmt.Println(a, b)
}

func split(sum int) (x,y int){
	x= sum * 4 /9
	y= sum-x
	return
}
func printsplit() {
	fmt.Println(split(17))
}


///

//package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int 
	speed        int
    distance     int
}
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    return Car{
        speed: speed,
        batteryDrain : batteryDrain,
        battery: 100, 
        distance: 0,
    }
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	panic("Please implement the CanFinish function")
}
