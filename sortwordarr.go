package piscine

func SortWordArr(array []string) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if comparearrayascii([]byte(array[j]), []byte(array[i])) {
				// Swap the elements
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func comparearrayascii(a, b []byte) bool {
	// Compare byte slices lexicographically
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return true // a is less than b
		} else if a[i] > b[i] {
			return false // a is greater than b
		}
	}
	// If we reach here, it means one string is a prefix of the other
	return len(a) < len(b) // shorter string is less
}
