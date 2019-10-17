package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	SetNbr(n)
}
func SetNbr(n int) {
	m := '0'
	if n == 0 {
		z01.PrintRune(m)
		return
	}
	for c := 1; c <= n%10; c++ {
		m++
	}
	for c := -1; c >= n%10; c-- {
		m++
	}
	if n/10 != 0 {
		SetNbr(n / 10)
	}
	z01.PrintRune(m)
	return
}
