package piscine

func Atoi(str string) int {
	//strReversed := StrRev(str)
	x := 0
	mm := false
	if StrLen(str) != 0 {
		if rune(str[0]) == '-' {
			for _, y := range []rune(str) {
				if y == '-' {
					if mm {
						return 0
					}
					mm = true
					continue
				} else if y <= '9' && y >= '0' {
					d := '0'
					j := 0
					for ; d != y; d++ {
						j--
					}
					if x != 0 || j != 0 {
						x = x*10 + j
					}
				} else {
					return 0
				}
			}
			return x
		} else {
			for _, y := range []rune(str) {
				if y <= '9' && y >= '0' {
					d := '0'
					j := 0
					for ; d != y; d++ {
						j++
					}
					if x != 0 || j != 0 {
						x = x*10 + j
					}
				} else {
					return 0
				}
			}
			return x
		}
	} else {
		return 0
	}
}