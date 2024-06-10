package main

import (
	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	q := n
	for _, v := range a {
		q = f(q, v)
	}
	r := Itoa(q)
	for _, char := range r {
		z01.PrintRune(char)
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
		digits := n % 10
		q = string(rune('0'+digits)) + q
		n /= 10
	}
	return sign + q
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

// package main
func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	z01.PrintRune('\n')

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
