package main

import (
	"os"
	piscine ".."
)

func main() {
	x := os.Args
	for i, y := range x {
		if i != 0 {
			piscine.PrintStr(y)
		}
	}
}	