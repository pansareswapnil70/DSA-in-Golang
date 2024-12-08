package main

import (
	"fmt"
)

func main() {
	fmt.Println("Min cost of climbing stairs:")
	arr := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	n := len(arr)
	dpStairs := make([]int, n+1)
	minCost := minCostOfClimbingStairs(arr, n, dpStairs)
	fmt.Print("Min cost of climbing stairs is: ", minCost)
}

func minCostOfClimbingStairs(arr []int, n int, dpStairs []int) int {
	if dpStairs[n] != 0 {
		return dpStairs[n]
	}
	if n <= 1 {
		dpStairs[n] = 0
		return dpStairs[n]
	}
	dpStairs[n-1] = minCostOfClimbingStairs(arr, n-1, dpStairs)
	dpStairs[n-2] = minCostOfClimbingStairs(arr, n-2, dpStairs)
	return min(dpStairs[n-1]+arr[n-1], dpStairs[n-2]+arr[n-2])
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
