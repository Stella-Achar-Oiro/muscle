package main

import (
	"github.com/01-edu/z01"
)

func PrintBits(octe byte) {
	for i := 7; i >= 0; i-- {
		if octe&(1<<i) != 0 {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintBits(2) // Example usage
}
