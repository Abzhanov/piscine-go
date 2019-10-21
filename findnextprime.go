package piscine

func FindNextPrime(nb int) int {
	for true {
		if IsPrime(nb) {
			return nb
		} else {
			nb++
		}
	}
	return 0
}
