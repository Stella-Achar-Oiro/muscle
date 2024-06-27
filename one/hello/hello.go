package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

// func main() {
// 	PrintStr("Hello World!")
// }

// func PrintStr(s string) {
// 	for _, r := range s {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	s := "Hello World!"
// 	for i := 0; i <= len(s)-1; i++ {
// 		z01.PrintRune(rune(s[i]))
// 	}
// 	z01.PrintRune('\n')
// }

func main() {
	s := "Hello World!"
	for i, v := range s {
		fmt.Print(i, string(v))
	}
	z01.PrintRune('\n')
}
