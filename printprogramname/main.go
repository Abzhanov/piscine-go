package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	x := os.Args
	res := x[0]
	for _, s := range res {
		z01.PrintRune(s)
	}
	z01.PrintRune(10)
}
