package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, y := range s {
		z := 0
		for i := '1'; i <= y; i++ {
			z++
		}
		x = x*10 + z
	}
	return x
}

func check(s string) bool {
	for _, y := range s {
		if y == ' ' || y < '0' || y > '9' {
			return false
		}
	}
	return true
}
