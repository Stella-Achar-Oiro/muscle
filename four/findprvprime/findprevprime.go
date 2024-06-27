package main

import "fmt"

func FindPrevPrime(n int) int {
	if n < 2 {
		return 0
	}
	if Isprime(n) {
		return n
	}
	return FindPrevPrime(n - 1)
}

func Isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(10))
	fmt.Println(FindPrevPrime(9))
	fmt.Println(FindPrevPrime(8))
	fmt.Println(FindPrevPrime(7))
	fmt.Println(FindPrevPrime(6))
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
	fmt.Println(FindPrevPrime(3))
	fmt.Println(FindPrevPrime(2))
	fmt.Println(FindPrevPrime(1))
}
