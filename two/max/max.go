package main

import "fmt"

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	q := a[0]
	for _, r := range a {
		if r > q {
			q = r
		}
	}
	return q
}

func main() {
	a := Max([]int{24, 56, 78, 98, 909})
	fmt.Println(a)
}
