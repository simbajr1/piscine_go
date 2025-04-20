package piscine

func Capitalize(s string) string {
	newWord := true
	result := ""
	for _, char := range s {
		isAlpha := (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9')

		if isAlpha {
			if newWord {
				if char >= 'a' && char <= 'z' {
					char = char - 32
				}
				newWord = false
			} else {
				if char >= 'A' && char <= 'Z' {
					char = char + 32
				}
			}
		} else {
			newWord = true
		}
		result += string(char)
	}
	return result
}
