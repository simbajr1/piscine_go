package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == ' ' || ch == '\t' || ch == '\n' {
			if len(word) > 0 {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	if len(word) > 0 {
		result = append(result, word)
	}

	return result
}
