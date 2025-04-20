package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 6 * 7 // 42
	ptr.y = 7 * 3 // 21
}

func PrintNumber(n int) {
	ne := '0'
	for i := 0; i < n; i++ {
		ne++
	}
	z01.PrintRune(ne)
}

func main() {
	points := &point{}
	setPoint(points)

	for _, r := range "x = " {
		z01.PrintRune(r)
	}
	PrintNumber(points.x / (5 * 2))
	PrintNumber(points.x % (5 * 2))

	for _, r := range ", y = " {
		z01.PrintRune(r)
	}
	PrintNumber(points.y / (5 * 2))
	PrintNumber(points.y % (5 * 2))
	z01.PrintRune('\n')
}
