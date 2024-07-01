package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) < 2 {
// 		return
// 	}
// 	args := os.Args[1:]

// 	var al string

// 	for _, arg := range args {
// 		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
// 			al = "Allert!!!"
// 		}
// 	}

// 	for _, s := range al {
// 		z01.PrintRune(s)
// 	}
// 	z01.PrintRune('\n')

// }

func main() {
	comcheck := map[string]bool{
		"01":        true,
		"galaxy":    true,
		"galaxy 01": true,
	}
	for _, r := range os.Args[1:] {
		if comcheck[r] {
			PrintStr("Alert!!!")
		}
	}
}

func PrintStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')

}
