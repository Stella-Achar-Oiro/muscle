package main

import "fmt"

func Rot13(s string) string {
	q := []rune(s)
	for i, v := range s {
		if v >= 'a' && v <= 'z' {
			q[i] = 'a' + (v-'a'+13)%26
		} else if v >= 'A' && v <= 'Z' {
			q[i] = 'A' + (v-'A'+13)%26
		}
	}
	return string(q)
}

func main() {
	fmt.Println(Rot13("abckzv"))
}
