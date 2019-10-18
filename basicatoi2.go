package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, y := range s {
		if y > '9' && y < '0' {
			return 0
		} else {
		z := 0 
		for i := '1'; i <= y; i++ {
			z++
		}
		x = x*10 + z
		}	
	}		
	return x
}
