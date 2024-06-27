package main

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println(slice)
		return
	}
	count, newSlice, outputSlice := 0, []int{}, [][]int{}

	for _, r := range slice {
		newSlice = append(newSlice, r)
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
	fmt.Println(outputSlice)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
