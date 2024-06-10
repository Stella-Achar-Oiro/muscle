package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	arg1, arg2 := os.Args[1], os.Args[2]

	num1 := Atoi(arg1)
	num2 := Atoi(arg2)

	gcd := GCD(num1, num2)

	number := Itoa(gcd)
	for _, c := range number {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	var number int
	sign := 1

	for idx, char := range s {
		if char == '-' && idx == 0 {
			sign = -1
		} else if char == '+' && idx == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return number * sign
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
