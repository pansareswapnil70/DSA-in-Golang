package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array Subset Sum:")
	arr := []int{3, 34, 4, 12, 5, 2}
	sum := 77
	n := len(arr)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
		for j := 0; j <= sum; j++ {
			dp[i][j] = false // Initialize all values as false
		}
	}

	// Base cases: dp[0][i] is false because no sum is possible with an empty set of elements
	for i := 0; i <= sum; i++ {
		dp[0][i] = false // no subset possible for sum > 0 with 0 elements
	}

	// Base case: dp[i][0] is true because a sum of 0 is always possible (using an empty subset)
	for i := 0; i <= n; i++ {
		dp[i][0] = true //subset sum of 0 is always achievable
	}
	boolResult := isSubsetSum(arr, sum, n, dp)
	fmt.Println("Subset sum is", boolResult)
}

func isSubsetSum(arr []int, sum int, n int, dp [][]bool) bool {
	if dp[n][sum] != false {
		return dp[n][sum]
	}
	if n == 0 {
		dp[n][sum] = false
		return dp[n][sum]
	}
	if sum == 0 {
		dp[n][sum] = true
		return dp[n][sum]
	}

	if arr[n-1] <= sum {
		dp[n][sum] = isSubsetSum(arr, sum-arr[n-1], n-1, dp) || isSubsetSum(arr, sum, n-1, dp)
	} else {
		dp[n][sum] = isSubsetSum(arr, sum, n-1, dp)
	}
	return dp[n][sum]
}
