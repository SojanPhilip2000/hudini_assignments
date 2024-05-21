package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {
	// TODO: Implement Calculator function
	result := 0.0
	if operand == "+" {
		result = value1 + value2
	} else if operand == "-" {
		result = value1 - value2
	} else if operand == "*" {
		result = value1 * value2
	} else if operand == "/" {
		if value2 == 0 {
			errors.New("cant dvide by zero ")
		} else {
			result = value1 / value2
		}
	}
	return result, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 10 + 20): ")
	input, err := reader.ReadString('\n')
	new := strings.Split(strings.TrimSpace(input), " ")
	length := len(new)
	if length != 3{
		fmt.Printf("Error reading input")

	}
	value1, err := strconv.ParseFloat(new[0], 64)
	value2, err := strconv.ParseFloat(new[2], 64)
	operand := new[1]
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	} else {
		result, err := Calculator(value1, value2, operand)

		if err != nil {
			fmt.Println("experrion is not correct")
		}
		fmt.Printf("Result: %v\n", result)

	}

	//input = strings.TrimSpace(input)
	//parts := strings.Split(input, " ")

	// TODO: Check if exact 3 parts are given. If not, throw an error

	// TODO: Take all 3 values and pass to calculator function

	// TODO: Print results
}
