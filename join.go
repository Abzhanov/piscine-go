package piscine

func Join(strs []string, sep string) string {
	res := ""
	//l := StrLen(strs)
	m := 0
	for range strs {
		m++
	}
	for i := 0; i < m; i++ {
		if m-i == 1 {
			res += strs[i]
		} else {
			res = res + strs[i] + sep
		}
	}
	return res
}
