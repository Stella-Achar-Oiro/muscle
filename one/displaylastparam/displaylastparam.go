package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <=1 {
		return
	}
	lastparam := os.Args[len(os.Args)-1]
	for _, r := range lastparam {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
