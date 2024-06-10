package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1, str2, i := os.Args[1], os.Args[2], 0

	

	for _, char := range str2 {
		if i < len(str1) && char == rune(str1[i]) {
			i++
		}
	}

	if i == len(str1) {
		for _, char := range str1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
