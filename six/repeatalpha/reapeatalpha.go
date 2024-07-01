package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	alpha := 0
	s := []rune(os.Args[1])
	if len(s) == 0 {
		return
	}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			alpha = int(r - 'a')
		} else if r >= 'A' && r <= 'Z' {
			alpha = int(r - 'A')
		}
		for i := 0; i <= alpha; i++ {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}