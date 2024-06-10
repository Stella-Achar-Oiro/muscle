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
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			alpha = int(char - 'A')
		} else if char >= 'a' && char <= 'z' {
			alpha = int(char - 'a')
		}
		for i := 0; i <= alpha; i++ {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
