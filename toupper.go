package piscine

func ToUpper(s string) string {
	var UpperChar string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			UpperChar += string(char - 32)
		} else {
			UpperChar += string(char)
		}
	}
	return UpperChar
}
