package piscine

func Sqrt(nb int) int {
	if nb < 1 {
		return 0
	}

	for x := 1; x*x <= nb; x++ {
		if x*x == nb {
			return x
		}
	}

	return 0
}
