package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	lastword, start, args := "", false, os.Args[1]
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] != ' ' {
			start = true
			lastword = string(args[i]) + lastword
		} else if start {
			break
		}
	}
	if len(lastword) > 0 {
		PrintStr(lastword)
	}
}
func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
