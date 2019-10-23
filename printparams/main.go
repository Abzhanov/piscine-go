package main

import (
	"github.com/01-edu/z01"
	"os"
)

func PrintStr(str string) {
	for _, h := range str {
		z01.PrintRune(h)
	}
	z01.PrintRune(10)
}

func main() {
	x := os.Args
	for i, y := range x {
		if i != 0 {
			PrintStr(y)
		}
	}
}
