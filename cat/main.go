package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r) // Print each rune
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return // No arguments, do nothing
	} else {
		for _, s := range args {
			file, err := os.Open(s)
			if err != nil {
				printString(err.Error() + "\n") // Print error message
				break
			} else {
				// Read the file content byte by byte
				var data []byte
				buf := make([]byte, 1) // Buffer to read one byte at a time
				for {
					n, err := file.Read(buf)
					if n == 0 {
						break // End of file
					}
					if err != nil {
						printString(err.Error() + "\n") // Print error message
						break
					}
					data = append(data, buf[0]) // Append the byte to the data slice
				}
				printString(string(data)) // Print file content
				file.Close()              // Close the file after reading
			}
		}
	}
}
