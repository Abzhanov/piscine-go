package piscine

func AppendRange(min, max int) []int {
	var res []int
	for i := min - 1; i < max-1; i++ {
		res = append(res, i+1)
	}
	return res
}
