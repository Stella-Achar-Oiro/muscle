package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	s, old, new := os.Args[1], os.Args[2], os.Args[3]

	replaced := replaceAll(s, old, new)
	PrintStr(replaced)
}

func replaceAll(s, old, new string) string {
	if len(s) < len(old) {
		return s
	}
	if s[:len(old)] == old {
		return new + replaceAll(s[len(old):], old, new)
	}
	return string(s[0]) + replaceAll(s[len(old):], old, new)
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
