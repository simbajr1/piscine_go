package piscine

func Split(s, sep string) []string {
	if sep == "" {
		var result []string
		for _, ch := range s {
			result = append(result, string(ch))
		}
		return result
	}

	var result []string
	start := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
			i += len(sep) - 1
		}
	}

	result = append(result, s[start:])
	return result
}
