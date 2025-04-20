package main

import (
	"os"
)

func printRune(r rune) {
	// Prints a single rune as a character to stdout
	os.Stdout.Write([]byte(string(r)))
}

func printString(s string) {
	// Prints the entire string character by character followed by a newline
	for _, r := range s {
		printRune(r)
	}
	os.Stdout.Write([]byte("\n"))
}

// stringToInt manually converts a string to an integer
func stringToInt(s string) (int64, bool) {
	var result int64
	var sign int64 = 1

	for i, r := range s {
		// Handle negative numbers
		if i == 0 && r == '-' {
			sign = -1
			continue
		}
		// Check if the character is a valid digit
		if r < '0' || r > '9' {
			return 0, false
		}
		// Check for overflow
		if result > (9223372036854775807-int64(r-'0'))/10 {
			return 0, false // overflow detected
		}
		result = result*10 + int64(r-'0')
	}
	return result * sign, true
}

// intToString converts an integer to a string manually
func intToString(num int64) string {
	if num == 0 {
		return "0"
	}

	str := ""
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}
	for num > 0 {
		str = string(num%10+'0') + str
		num /= 10
	}
	return sign + str
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	// Convert values to integers (using int64 to avoid overflow)
	val1, valid1 := stringToInt(os.Args[1])
	val2, valid2 := stringToInt(os.Args[3])

	// Check if the values are valid and within the 64-bit integer range
	if !valid1 || !valid2 {
		return // invalid input
	}

	operator := os.Args[2]

	var result string
	switch operator {
	case "+":
		// Overflow check for addition
		if val1 > 9223372036854775807-val2 || val1 < -9223372036854775808+val2 {
			return // overflow detected, print nothing
		}
		result = intToString(val1 + val2)
	case "-":
		// Overflow check for subtraction
		if (val1 < 0 && val2 > 0 && val1 < -9223372036854775808+val2) || (val1 > 0 && val2 < 0 && val1 > 9223372036854775807+val2) {
			return // overflow detected, print nothing
		}
		result = intToString(val1 - val2)
	case "*":
		// Overflow check for multiplication
		if val1 > 9223372036854775807/val2 || val1 < -9223372036854775808/val2 {
			return // overflow detected, print nothing
		}
		result = intToString(val1 * val2)
	case "/":
		if val2 == 0 {
			printString("No division by 0")
			return
		}
		result = intToString(val1 / val2)
	case "%":
		if val2 == 0 {
			printString("No modulo by 0")
			return
		}
		result = intToString(val1 % val2)
	}

	// Only print result if there was no error
	if result != "" {
		printString(result)
	}
}
