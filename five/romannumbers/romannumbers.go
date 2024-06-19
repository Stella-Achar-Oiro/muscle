package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		printError()
		return
	}

	arg := os.Args[1]
	num := Atoi(arg)
	if num <= 0 || num >= 4000 {
		printError()
		return
	}

	roman, calculation := toRoman(num)
	printString(calculation)
	printString(roman)
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

func toRoman(num int) (string, string) {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	calculation := ""
	for i := 0; i < len(value); i++ {
		for num >= value[i] {
			num -= value[i]
			roman += symbol[i]
			if calculation != "" {
				calculation += "+"
			}
			calculation += symbol[i]
		}
	}
	return roman, calculation
}

func printString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func printError() {
	msg := "ERROR: cannot convert to roman digit"
	for _, c := range msg {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
