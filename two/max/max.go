package main

import "fmt"

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	q := a[0]
	for _, char := range a {
		if char > q {
			q = char
		}
	}
	return q
}
func main() {
	a := []int{23, 32, 54, 67, 102, 231}
	max := Max(a)
	fmt.Println(max)
}
