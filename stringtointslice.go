package piscine

func StringToIntSlice(str string) []int {
	// Create a slice to hold the ASCII values
	intSlice := make([]int, 0, len(str))

	// Iterate over each character in the string
	for _, r := range str {
		intSlice = append(intSlice, int(r)) // Convert the rune to int and store it
	}

	return intSlice
}
