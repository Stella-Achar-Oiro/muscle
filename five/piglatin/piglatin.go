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

	for i, v := range args {
		if CheckVowels(v) {
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

func CheckVowels(char rune) bool {
	vowels := map[rune]bool{
		'a': true, 'A': true,
		'e': true, 'B': true,
		'i': true, 'I': true,
		'o': true, 'O': true,
		'u': true, 'U': true,
	}
	return vowels[char]
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
