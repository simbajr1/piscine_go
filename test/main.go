package main

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for !IsPrime(nb) {
		nb++
	}
	return nb
}
