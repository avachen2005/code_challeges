package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{}

	fmt.Printf("\ninput %d, %v\n=====\n", 5, searchMatrix(matrix, 5))

	matrix = [][]int{
		{},
	}

	fmt.Printf("\ninput %d, %v\n=====\n", 5, searchMatrix(matrix, 5))

	matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	fmt.Printf("\ninput %d, %v\n=====\n", 5, searchMatrix(matrix, 5))
	fmt.Printf("\ninput %d, %v\n", 20, searchMatrix(matrix, 20))

	matrix = [][]int{
		{1, 4, 7, 11, 15},
	}

	fmt.Printf("\ninput %d, %v\n=====\n", 7, searchMatrix(matrix, 7))
	fmt.Printf("\ninput %d, %v\n=====\n", 4, searchMatrix(matrix, 4))
	fmt.Printf("\ninput %d, %v\n", 20, searchMatrix(matrix, 20))
}

func searchMatrix(matrix [][]int, target int) bool {

	for _, e1 := range matrix {
		for _, e2 := range e1 {
			if target == e2 {
				return true
			}
		}
	}

	return false
}
