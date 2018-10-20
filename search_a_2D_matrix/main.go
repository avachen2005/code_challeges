package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	fmt.Printf("input %d, %v\n=====\n", 5, searchMatrix(matrix, 5))
	fmt.Printf("input %d, %v\n", 20, searchMatrix(matrix, 20))
}

func searchMatrix(matrix [][]int, target int) bool {

	x := int(len(matrix[0]) / 2)
	y := int(len(matrix) / 2)

	if len(matrix) == 0 {
		return false
	}

	var from_x, from_y, to_x, to_y int

	if matrix[x][y] > target {

		from_x = 0
		from_y = 0
		to_x = x + 1
		to_y = y + 1

		fmt.Printf("from_x: %d, from_y: %d, to_x: %d, to_y: %d\n", from_x, from_y, to_x, to_y)
	} else if matrix[x][y] == target {

		return true

	} else {

		from_x = x
		from_y = y
		to_x = len(matrix[0])
		to_y = len(matrix)

	}

	i := from_x
	j := from_y

	for i < to_x {

		j = from_y

		for j < to_y {

			fmt.Printf("(%d,%d,%d) ", i, j, matrix[i][j])
			if matrix[i][j] == target {
				return true
			}

			j = j + 1
		}
		i = i + 1
		fmt.Printf("\n")
	}

	return false
}
