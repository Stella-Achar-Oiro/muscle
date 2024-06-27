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

	num := Atoi(os.Args[1])
	if num == 0 && os.Args[1] != "0" {
		PrintError()
		return
	}

	PrintHex(num)
}

func Atoi(s string) int {
	num := 0

	for _, v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		num = num*10 + int(v-'0')
	}
	return num
}

func PrintHex(num int) {
	if num == 0 {
		PrintStr("0")
		return
	}

	hexaDecimal := "0123456789abcdef"
	hexStr := ""
	for num > 0 {
		hexStr = string(hexaDecimal[num%16]) + hexStr
		num /= 16
	}
	PrintStr(hexStr)
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func PrintError() {
	PrintStr("ERROR")
}
