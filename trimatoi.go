package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	foundDigit := false

	for _, char := range s {
		if char == '-' && !foundDigit {
			sign = -1
			continue
		}
		if char == '+' && !foundDigit {
			continue
		}
		if char >= '0' && char <= '9' {
			foundDigit = true
			result = result*10 + int(char-'0')
		}
	}
	return result * sign
}
