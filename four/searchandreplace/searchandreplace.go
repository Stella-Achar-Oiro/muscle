package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str, old, new := os.Args[1], os.Args[2], os.Args[3]

	result := Replace(str, old, new)
	PrintStr(result)
}

func Replace(s, old, new string) string {
	if len(s) < len(old) {
		return s
	}
	if s[:len(old)] == old {
		return new + Replace(s[len(old):], old, new)
	}
	return string(s[0]) + Replace(s[1:], old, new)
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
