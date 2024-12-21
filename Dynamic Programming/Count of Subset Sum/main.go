package main

import (
	"fmt"
)

func main() {
	fmt.Println("Count of Subsets of given sum:")
	nums := []int{1, 2, 3, 3}
	sum := 0
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			dp[i][j] = -1
		}
	}
	count := countOfSubset(nums, n, sum, dp)
	fmt.Println("Count of subsets are:", count)
}

func countOfSubset(nums []int, n int, sum int, dp [][]int) int {
	if dp[n][sum] != -1 {
		return dp[n][sum]
	}
	if sum == 0 {
		dp[n][sum] = 1
		return dp[n][sum]
	}
	if n == 0 {
		dp[n][sum] = 0
		return dp[n][sum]
	}

	if nums[n-1] <= sum {
		dp[n][sum] = countOfSubset(nums, n-1, sum-nums[n-1], dp) + countOfSubset(nums, n-1, sum, dp)
	} else {
		dp[n][sum] = countOfSubset(nums, n-1, sum, dp)
	}
	return dp[n][sum]
}
