package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Start from index 1 to skip the program name
	args := os.Args[1:]

	// Iterate through each argument
	for _, arg := range args {
		// Print each character of the argument
		for _, char := range arg {
			z01.PrintRune(char)
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
