// package main

// import (
// 	"github.com/01-edu/z01"
// )

// func PrintBits(octe byte) {
// 	for i := 7; i >= 0; i-- {
// 		if octe&(1<<i) != 0 {
// 			z01.PrintRune('1')
// 		} else {
// 			z01.PrintRune('0')
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	PrintBits(2) // Example usage
// }

package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	arg := os.Args[1]
	num := 0
	for _, c := range arg {
		if c < '0' || c > '9' {
			print("00000000")
			return
		}
		num = num*10 + int(c-'0')
	}

	binary := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		bit := byte((num >> uint(i)) & 1)
		binary[7-i] = bit + '0'
	}
	print(string(binary))
}
