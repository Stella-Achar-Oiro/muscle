// Seth Solution

package main

import "fmt"

// func FindPrevPrime(nb int) int {
// 	if nb <= 1 {
// 		return 0
// 	}
// 	var prev = 0
// 	for i := 2; i <= nb; i++ {
// 		if IsPrime(i) {
// 			prev = i
// 		}
// 	}

// 	return prev
// }

// func IsPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}

// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}

// 	}
// 	return true
// }

//func main() {
	//fmt.Println(FindPrevPrime(5))
	//fmt.Println(FindPrevPrime(4))
//}

// Stella Solution

//package main

//import "fmt"

// func FindPrevPrime(nb int) int {
// 	if nb < 2 {
// 		return 0
// 	}
// 	if IsPrime(nb) {
// 		return nb
// 	}
// 	return FindPrevPrime(nb - 1)
// }

// func IsPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println(FindPrevPrime(5))
// 	fmt.Println(FindPrevPrime(4))
// }

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}


func FindPrevPrime(num int) int {
	if num < 2 {
		return 0
	}
	if IsPrime(num) {
		return(num)
	}
	return FindPrevPrime(num-1)
}
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i:= 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}