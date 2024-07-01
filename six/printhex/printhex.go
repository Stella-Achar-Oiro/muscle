package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		PrintStr("ERROR")
		return
	}
	input := os.Args[1]
	n := Atoi(input)
	if n == 0 && input != "0" {
		PrintStr("ERROR")
		return
	}

	PrintHex(n)
}

func PrintHex(n int) {
	if n == 0 {
		PrintStr("0")
		return
	}

	hex := "0123456789abcdef"
	q := ""
	for n > 0 {
		q = string(hex[n%16]) + q
		n /= 16
	}
	PrintStr(q)
}

func Atoi(s string) int {
	q := 0

	for _, v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		q = q*10 + int(v-'0')
	}
	return q
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
