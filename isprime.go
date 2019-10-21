package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 1; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
