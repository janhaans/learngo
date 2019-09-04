package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Ask for input
	fmt.Print("Enter a grade [0-100]: ")
	// Create a reader to read from terminal
	reader := bufio.NewReader(os.Stdin)
	// Read input until new line
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Remove the new line character from string
	input = strings.TrimSpace(input)
	// Converse value of input from type string to float64
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	if grade > 55 {
		fmt.Println("You have passed the test")
	} else {
		fmt.Println("You have not passed the test")
	}
}
