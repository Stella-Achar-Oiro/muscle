package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := len(os.Args[1:])
	paramcount := Itoa(count)
	PrintStr(paramcount)
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	q := ""
	for n > 0 {
		num := n % 10
		q = string(rune('0'+num)) + q
		n /= 10
	}
	return sign + q
}
