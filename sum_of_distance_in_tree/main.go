package main

import (
	"fmt"
)

func main() {

}

func sumOfDistancesInTree(N int, edges [][]int) []int {

	result := make([]int, N)
	/*
	   N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
	*/
	for i := 0; i < N; i++ {

		val := 0
		cnt := 0
		layer := 1

		for cnt < len(edges) {
			from := edges[cnt][0]
			to := edges[cnt][1]

			if from == i {
				val = val + layer
			} else {
				layer = layer + 1
			}
		}

		result[i] = val
	}
}
