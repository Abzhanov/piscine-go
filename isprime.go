package piscine

func IsPrime(nb int) bool {
	if nb<=1 {
		return false
	}
	for i := 0; i < nb; i++{
		if nb % 2 == 0 {
			return false
		}		
	}
	return true
}
