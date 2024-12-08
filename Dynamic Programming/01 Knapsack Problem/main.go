package main

import (
	"fmt"
)

func main() {
	fmt.Println("0/1 Knapsack Problem")
	N := 5                                  // number of items
	W := 10                                 // maximum weight bag can handle
	profits := []int{60, 100, 120, 80, 150} // profit associated with items
	weights := []int{10, 20, 30, 20, 50}    // weight associated with items
	result := knapSack(profits, weights, W, N)
	fmt.Println("Maximum profit obtained is: ", result)
}

func knapSack(profits []int, weights []int, W int, N int) int {
	if N == 0 || W == 0 {
		return 0
	}

	if weights[N-1] <= W {
		return max(profits[N-1]+knapSack(profits, weights, W-weights[N-1], N-1),
			knapSack(profits, weights, W, N-1))
	} else {
		return knapSack(profits, weights, W, N-1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
