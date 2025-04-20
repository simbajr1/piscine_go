package piscine

func DescendAppendRange(max, min int) []int {
	// If max is less than or equal to min, return an empty slice
	if max <= min {
		return []int{}
	}

	// Create a slice to store the result
	var result []int

	// Loop from max down to min (exclusive)
	for i := max; i > min; i-- {
		result = append(result, i)
	}

	return result
}
