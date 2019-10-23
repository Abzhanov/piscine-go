package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	x := os.Args
	for j := 0; j < len(x)-1; j++ {
		for i := 0; i < len(x)-1; i++ {
			if x[i] > x[i+1] {
				temp := x[i]
				x[i] = x[i+1]
				x[i+1] = temp
			}
		}
	}
	for index, item := range x {
		if index != 0 {
			for _, y := range item {
				z01.PrintRune(y)
			}
			z01.PrintRune(10)
		}
	}
}
