package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var res []int
	if n < 0 {
		z01.PrintRune('-')
	} else if n == 0 {
		z01.PrintRune('0')
	}
	j := 0
	for i := n; i > 0; i /= 10 {
		res = append(res, i%10)
		j++
	}
	for i := 0; i <= j-1; i++ {
		for k := i + 1; k <= j-1; k++ {
			if res[i] > res[k] {
				tr := res[i]
				res[i] = res[k]
				res[k] = tr
			}
		}
	}
	for i := 0; i <= j-1; i++ {
		z01.PrintRune(48 + rune(res[i]))
	}
}
