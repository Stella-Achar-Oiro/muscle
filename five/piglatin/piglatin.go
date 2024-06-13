package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args, word, foundVowel := os.Args[1], "", false

	for i, char := range args {
		if checkVowels(char) {
			if i == 0 {
				word = args + "ay"
			} else {
				word = args[i:] + args[:i] + "ay"
			}
			foundVowel = true
			break
		}
	}

	if !foundVowel {
		word = "No vowels"
	}
	PrintStr(word)
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func checkVowels(c rune) bool {
	return c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I' || c == 'o' || c == 'O' || c == 'u' || c == 'U'
}
