package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	n := BasicAtoi(os.Args[1])
	IsPowerOf2 := (n > 0) && (n&(n-1) == 0)

	if IsPowerOf2 {
		PrintStr("true")
	} else {
		PrintStr("false")
	}
}

func BasicAtoi(s string) int {
	q := 0
	for _, r := range s {
		q = q*10 + int(r-'0')
	}
	return q
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
