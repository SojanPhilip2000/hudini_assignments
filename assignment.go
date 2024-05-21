package main

import {
	"fmt"
	"math"
}

// 1. Calculate Average
// Objective: Write a function that takes an array of numbers and returns the average. Use loops and basic arithmetic.
// Function signature:

func calculateAverage(numbers []int) int {
	var sum = 0
	var i = 0
	for i = 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum / len(numbers)
}

// 2. Check Age
// Write a function that takes an age and prints whether the person is a minor, a young adult, or an adult.
// Use conditional statements.
// Function signature:
func checkAge(age int) string {
	if age > 150 {
		fmt.Println("age limit exceeded")
	} else if age < 18 {
		fmt.Println("minor")
	} else if age < 25 {
		fmt.Println("young adult")
	} else {
		fmt.Println("adult")
	}
	return ""
}

// 3. Reverse String
// Objective: Create a function that reverses a string. This will demonstrate basic string manipulation and for loops.
// Function signature:
// func reverseString(str string) string {
// 	// Write your code here to reverse and return the string.
// 	var reverse = ""
// 	var i = len(str)
// 	for i = len(str) - 1; i >= 0; i-- {
// 		reverse += string(str[i])
// 	}
// 	return reverse
// }

// func reverseString(word string) string{
// 	var reversedWord = ""
// 	var i=len(word)
// 	for i=len(word)-1; i>=0; i--{

// 	}
// }
// 4. Find Largest Number
// Objective: Write a function that takes an array of numbers and returns the largest number.
// Use loops and conditional statements to solve the problem.
// Function signature:
func findLargestNumber(numbers []int) int{
    // Write your code here to find and return the largest number in the array.
	var largest = 0
	var i =0
	for i =0; i<=len(numbers)-1; i++{
		if numbers[i]> largest{
			largest =numbers[i]
		}
	}
	return largest
}

func main() {
	var slice = []int{10, 20, 30, 40}
	fmt.Println(calculateAverage(slice))
	checkAge(23)
	//fmt.Println(reverseString("string1"))
	fmt.Println(findLargestNumber(slice))
	var b= Counter{count :0}
	b=b.add()
	fmt.Println(b.count)
	monkeyCount(10)
}

type Counter struct{
	count int 
}
func (a Counter) add() Counter{
		a.count = a.count +1
		return a
}
func (a Counter) subtract() Counter{
		a.count =a.count - 1
		return a
}
