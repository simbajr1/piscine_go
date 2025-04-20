package piscine

func LastRune(s string) rune {
	sRune := []rune(s)
	return sRune[len(sRune)-1]
}
