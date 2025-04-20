package piscine

func Atoi(s string) int {
	myInt := 0
	sign := 1
	started := false

	for _, c := range s {
		if c == '-' {
			if started {
				return 0
			}
			sign = -1
			started = true
			continue
		} else if c == '+' {
			if started {
				return 0
			}
			started = true
			continue
		}
		if c < '0' || c > '9' {
			if started {
				return 0
			}
			continue
		}
		started = true
		myInt = myInt*10 + int(c-'0')
	}
	return myInt * sign
}
