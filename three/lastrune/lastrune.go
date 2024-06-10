package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	a := []rune(s)
	return a[len(a)-1]
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
