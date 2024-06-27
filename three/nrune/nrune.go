package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	if n <= 0 || n >= len(s) {
		return 0
	}
	return []rune(s)[n-1]
}

func main() {
	z01.PrintRune(NRune("Achar", 4))
	z01.PrintRune('\n')
}
