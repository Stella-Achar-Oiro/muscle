package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := Atoi(os.Args[1])
	PrintStr(PrintBits(input))
}

func PrintBits(num int) string {
	if num < 0 || num > 255 {
		return "00000000"
	}
	bits := ""
	for i := 7; i >= 0; i-- {
		bits += string('0' + byte((num>>i)&1))
	}
	return bits
}

func Atoi(s string) int {
	num := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		num = num*10 + int(c-'0')
	}
	return num
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
