package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	p1 := "x = "
	p2 := points.x
	p3 := ", y = "
	p4 := points.y

	for _, v := range p1 {
		z01.PrintRune(v)
	}

	z01.PrintRune(rune(p2) + 10)
	z01.PrintRune(rune(p2) + 8)

	for _, v := range p3 {
		z01.PrintRune(v)
	}

	z01.PrintRune(rune(p4) + 29)
	z01.PrintRune(rune(p4) + 28)

	z01.PrintRune('\n')
}
