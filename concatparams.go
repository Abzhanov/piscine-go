package piscine

func l(d []string) int {
	m := 0
	for range d {
		m++
	}
	return m
}

func ConcatParams(args []string) string {
	res := ""
	for i, j := range args {
		res += j
		if i != l(args)-1 {
			res += "\n"
		}
	}
	return res
}
