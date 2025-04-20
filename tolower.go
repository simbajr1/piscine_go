package piscine

func ToLower(s string) string {
	var lowerChar string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			lowerChar += string(char + 32)
		} else {
			lowerChar += string(char)
		}
	}
	return lowerChar
}
