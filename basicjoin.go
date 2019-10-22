package piscine

func BasicJoin(strs []string) string {
	var res string
	for _, str := range strs {
		res += str
	}
	return res
}
