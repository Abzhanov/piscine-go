package piscine

func BasicAtoi(s string) int {
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
