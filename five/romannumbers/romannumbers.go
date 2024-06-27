package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		PrintError()
		return
	}

	arg := os.Args[1]
	num := Atoi(arg)
	if num <= 0 || num >= 4000 {
		PrintError()
		return
	}

	roman, calculation := ToRoman(num)
	PrintStr(calculation)
	PrintStr(roman)
}

func Atoi(s string) int {
	var number int
	for _, char := range s {
		if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		} else {
			return -1
		}
	}
	return number
}

func ToRoman(num int) (string, string) {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	calculation := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += sym[i]
			if calculation != "" {
				calculation += "+"
			}
			calculation += sym[i]
		}
	}
	return roman, calculation
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func PrintError() {
	msg := "ERROR: cannot convert to roman digit"
	PrintStr(msg)
}
