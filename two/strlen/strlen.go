package main

import (
	"fmt"
)

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func main() {
	r := StrLen("Yes dear")
	fmt.Println(r)
}
