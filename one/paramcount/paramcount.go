package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := len(os.Args[1:])
	Paramcount := Itoa(count)
	PrintStr(Paramcount)
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
		n = -n
		sign = "-"
	}
	q := ""
	for n > 0 {
		digits := n % 10
		q = string(rune('0'+digits)) + q
		n /= 10
	}
	return sign + q
}
