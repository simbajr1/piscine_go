package piscine

func StrRev(s string) string {
	var result string
	for _, c := range s {
		result = string(c) + result
	}
	return result
}
