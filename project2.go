// Objective:Create a command-line interface (CLI) application in Go that allows users to input a block of text and then calculates the frequency of each word in the text. This project will help you understand and implement basic Go concepts like variables, loops, conditionals, maps, functions, and strings.
// Requirements:
// Input Text: Allow users to input a block of text.
// Process Text: Count the frequency of each word in the text.
// Display Frequencies: Display the word frequencies in a readable format.
// Exit: Allow the user to exit the application.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a text block: ")
	input, _ := reader.ReadString('\n')
	getCount(input)
	count(input)
}

func getCount(input string) {
	words := strings.Split(strings.TrimSpace(input), " ")
	result := make(map[string]int)

	for _, word := range words {
		if result[word] == 0 {//result ennu parayana map il engana word enna key vanne
			result[word] = 1
		} else {
			result[word] += 1
		}
	}

	for word, count := range result {
		fmt.Printf("Word: %s and the count: %d \n", word, count)
	}

}

func count(input string){
	new := strings.Split(strings.TrimSpace(input), " ")
	output := make(map[string]int)
	for _, v := range(new){
		output[v] += 1
	}
	fmt.Println(output) 
}


