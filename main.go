package main

import "fmt"

func main() {
	fmt.Println("sm...")
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	// line := make([]int, n)
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}
