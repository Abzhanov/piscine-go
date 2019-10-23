package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	x := os.Args
	i := 0
	for range x {
		i++
	}
	name := sortArray(x[1:i])
	for j := 0; j < (i - 1); j++ {
		for _, letter := range name[j] {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}
}

func sortArray(str []string) []string {
	j := 0
	name := str
	for range str {
		j++
	}
	for i := 0; i < j; i++ {
		for k := i + 1; k < j; k++ {
			if name[i] > name[k] {
				name[i], name[k] = name[k], name[i]
			}
		}
	}
	return name
}
