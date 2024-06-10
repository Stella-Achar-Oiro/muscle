package main

import "github.com/01-edu/z01"

func Chunk(slice []int, size int) {
	if size <= 0 {
		z01.PrintRune('\n')
		return
	}
	chunks := [][]int{}
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	PrintChunks(chunks)
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

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
