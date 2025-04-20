package main

import (
	"fmt"
	"os"
)

func main() {
	// Check the number of arguments
	if len(os.Args) < 2 {
		// Case: No arguments passed
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		// Case: Too many arguments
		fmt.Println("Too many arguments")
		return
	}

	// Read the file specified as argument
	fileName := os.Args[1]

	// Try to read the file content using os.ReadFile
	content, err := os.ReadFile(fileName)
	if err != nil {
		// Handle error if file reading fails
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the content of the file using fmt
	fmt.Print(string(content))
}
