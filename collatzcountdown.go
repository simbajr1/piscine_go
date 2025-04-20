package piscine

func CollatzCountdown(start int) int {
	// Check if the starting number is 0 or negative
	if start <= 0 {
		return -1
	}

	// Initialize step count
	steps := 0

	// Apply the Collatz steps until the number reaches 1
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = 3*start + 1
		}
		steps++
	}

	return steps
}
