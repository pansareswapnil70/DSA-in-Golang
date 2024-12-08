package main

import (
	"fmt"
)

func main() {
	fmt.Println("No. of ways to climb N stairs")
	stairsTestCasesSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	for _, n := range stairsTestCasesSlice {
		dpstairs := make([]int, n+1)
		ways := climbStairs(n, dpstairs)
		fmt.Println("No. of ways to climb ", n, " stairs is ", ways)
	}
}

func climbStairs(n int, dpstairs []int) int {
	if dpstairs[n] != 0 {
		return dpstairs[n]
	}
	if n <= 1 {
		dpstairs[n] = 1
		return dpstairs[n]
	}
	dpstairs[n] = climbStairs(n-1, dpstairs) + climbStairs(n-2, dpstairs)
	return dpstairs[n]
}
