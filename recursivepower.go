package piscine

func RecursivePower(nb int, power int) int {
	a := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		a = nb * RecursivePower(nb, power-1)
	}
	return a
}
