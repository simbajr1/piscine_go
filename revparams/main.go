package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get command-line arguments (excluding the program name)
	args := os.Args[1:]

	// Iterate through the arguments in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		// Print each argument
		for _, r := range args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
