package main

import (
	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	count := 0
	newSlice, outputSlice := []int{}, [][]int{}

	if size == 0 {
		return
	}
	if len(slice) == 0 {
		PrintEmpty()
		z01.PrintRune('\n')
		return
	}

	for _, n := range slice {
		newSlice = append(newSlice, n)
		count++
		if count == size {
			outputSlice = append(outputSlice, newSlice)
			newSlice = []int{}
			count = 0
		}
	}
	if len(newSlice) > 0 {
		outputSlice = append(outputSlice, newSlice)
	}

	PrintChunks(outputSlice)
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	q := ""
	for n > 0 {
		digits := n % 10
		q = string(rune('0'+digits)) + q
		n /= 10
	}
	return sign + q
}

func PrintChunks(chunks [][]int) {
	z01.PrintRune('[')
	for i, chunk := range chunks {
		if i > 0 {
			z01.PrintRune(' ')
		}
		z01.PrintRune('[')
		for j, v := range chunk {
			if j > 0 {
				z01.PrintRune(' ')
			}
			q := Itoa(v)
			for _, r := range q {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune(']')
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

func PrintEmpty() {
	z01.PrintRune('[')
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
