package piscine

func Abort(a, b, c, d, e int) int {
	// Declare an array with the five input values
	values := []int{a, b, c, d, e}

	// Implement the sorting logic manually to find the middle value
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				// Swap if the values are out of order
				values[i], values[j] = values[j], values[i]
			}
		}
	}

	// Return the middle element (the 3rd element)
	return values[2]
}
