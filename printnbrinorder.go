package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	counts := [10]int{}

	for n > 0 {
		digit := n % 10
		counts[digit]++
		n = n / 10
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < counts[i]; j++ {
			z01.PrintRune(rune('0' + i))
		}
	}
}
