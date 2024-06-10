// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	args, lastword, start := os.Args[1], "", false

// 	for i := len(args) - 1; i >= 0; i-- {
// 		if args[i] != ' ' {
// 			start = true
// 			lastword = string(args[i]) + lastword
// 		} else if start {
// 			break
// 		}
// 	}
// 	if len(lastword) > 0 {
// 		for _, char := range lastword {
// 			z01.PrintRune(char)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args, start, lastword := os.Args[1], false, ""

	for i := len(args) - 1; i >= 0; i-- {
		if args[i] != ' ' {
			start = true
			lastword = string(args[i]) + lastword
		} else if start {
			break
		}
	}
	if len(lastword) > 0 {
		for _, char := range lastword {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
