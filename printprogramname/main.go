package main

import (
	"fmt"
	"os"
)

func main() {
	programName := os.Args[0]

	for i := len(programName) - 1; i >= 0; i-- {
		if programName[i] == '\\' {
			programName = programName[i+1:]
			break
		}
	}
	for _, char := range programName {
		fmt.Print(string(char))
	}
	fmt.Println()
}
