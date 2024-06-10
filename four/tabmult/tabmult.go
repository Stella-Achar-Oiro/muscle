package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	q := 0
	for _, r := range s {
		q = q*10 + int(r-'0')
	}
	return q
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

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	num := BasicAtoi(args)

	for i := 1; i <= 9; i++ {
		table := num * i
		PrintStr(Itoa(i) + " x " + Itoa(num) + " = " + Itoa(table))
	}
}
