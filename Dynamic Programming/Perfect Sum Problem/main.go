package main

import (
	"fmt"
)

func main() {
	fmt.Println("Perfect Sum Problem")
	nums := []int{5, 2, 3, 10, 6, 8}
	n := len(nums)
	sum := 10
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			dp[i][j] = -1
		}
	}
	count := countSubsetSum(nums, n, sum, dp)
	fmt.Println("Count of subset sum is ", count)
}

func countSubsetSum(nums []int, n int, sum int, dp [][]int) int {
	if dp[n][sum] != -1 {
		return dp[n][sum]
	}
	if n == 0 && sum > 0 {
		dp[n][sum] = 0
		return dp[n][sum]
	} else if n == 1 {
		if sum == 0 {
			if nums[0] == 0 {
				dp[n][sum] = 2
				return dp[n][sum]
			}
			dp[n][sum] = 1
			return dp[n][sum]
		} else {
			if nums[0] == sum {
				dp[n][sum] = 1
				return dp[n][sum]
			} else {
				dp[n][sum] = 0
				return dp[n][sum]
			}

		}
	} else if n == 0 && sum == 0 {
		dp[n][sum] = 1
		return dp[n][sum]
	}
	if nums[n-1] <= sum {
		dp[n][sum] = countSubsetSum(nums, n-1, sum-nums[n-1], dp) + countSubsetSum(nums, n-1, sum, dp)
		return dp[n][sum]
	} else {
		dp[n][sum] = countSubsetSum(nums, n-1, sum, dp)
		return dp[n][sum]
	}
}
