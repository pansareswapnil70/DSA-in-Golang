package main

import (
	"fmt"
)

func main() {
	fmt.Println("Partition Subset Sum:")
	nums := []int{1, 5, 11, 5}
	n := len(nums)
	totalSum := 0
	for i := range nums {
		totalSum += nums[i]
	}

	if totalSum%2 != 0 {
		boolResult := false
		fmt.Println("Partition sum is: ", boolResult)
		return
	}

	sum := totalSum / 2
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}
	boolResult := partitionSum(nums, n, sum, dp)
	fmt.Println("Partition sum is: ", boolResult)
}

func partitionSum(nums []int, n int, sum int, dp [][]bool) bool {
	if n == 0 {
		dp[n][sum] = false
		return dp[n][sum]
	}
	if sum == 0 {
		dp[n][sum] = true
		return dp[n][sum]
	}

	if nums[n-1] <= sum {
		dp[n][sum] = partitionSum(nums, n-1, sum-nums[n-1], dp) || partitionSum(nums, n-1, sum, dp)
		return dp[n][sum]
	} else {
		dp[n][sum] = partitionSum(nums, n-1, sum, dp)
		return dp[n][sum]
	}
}
