package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if r >= 0 && r <= 31 {
			return false
		}
	}
	return true
}
