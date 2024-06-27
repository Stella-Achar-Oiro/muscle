package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	num, valid := Atoi(arg)
	if !valid {
		PrintBinaryString("00000000")
		return
	}

	binary := IntToBinary(num)
	PrintBinaryString(binary)
}

// Atoi converts a string to an integer without using strconv package.
func Atoi(s string) (int, bool) {
	var result int
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		result = result*10 + int(c-'0')
	}
	return result, true
}

// IntToBinary converts an integer to an 8-bit binary string.
func IntToBinary(n int) string {
	binary := make([]rune, 8)
	for i := 7; i >= 0; i-- {
		if n%2 == 1 {
			binary[i] = '1'
		} else {
			binary[i] = '0'
		}
		n /= 2
	}
	return string(binary)
}

// PrintBinaryString prints a string using z01.PrintRune.
func PrintBinaryString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
