package piscine

import (
	"github.com/01-edu/z01"
)

func JumpOver(str string) {
	if len(str) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Iterate over the string starting from index 2 and step by 3
	for i := 2; i < len(str); i += 3 {
		z01.PrintRune(rune(str[i]))
	}

	// Print a newline after printing the characters
	z01.PrintRune('\n')
}
