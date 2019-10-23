package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	x := os.Args
	y := x[0]
	for _, res := range y {
		z01.PrintRune(res)
	}
	z01.PrintRune(10)
}
