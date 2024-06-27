package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1, s2, i := os.Args[1], os.Args[2], 0

	for _, r := range s2 {
		if i < len(s1) && r == rune(s1[i]) {
			i++
		}
	}
	if i == len(s1) {
		PrintStr(s1)
	}
}

func PrintStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
