package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum < abs(target) || (sum+target)%2 != 0 {
		fmt.Println("No of ways are: ", 0)
		return
	}
	sum = (sum + target) / 2
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			dp[i][j] = -1
		}
	}
	ways := isSubsetSum(nums, n, sum, dp)
	fmt.Println("No of ways are: ", ways)
}

func isSubsetSum(nums []int, n int, sum int, dp [][]int) int {
	if dp[n][sum] != -1 {
		return dp[n][sum]
	}
	if n == 0 {
		dp[n][sum] = 0
		return dp[n][sum]
	} else if n == 1 {
		if sum == 0 {
			if nums[n-1] == 0 {
				dp[n][sum] = 2
				return dp[n][sum]
			}
			dp[n][sum] = 1
			return dp[n][sum]
		} else {
			if nums[n-1] == sum {
				dp[n][sum] = 1
				return dp[n][sum]
			}
			dp[n][sum] = 0
			return dp[n][sum]
		}
	} else if n == 0 && sum == 0 {
		dp[n][sum] = 1
		return dp[n][sum]
	}

	if nums[n-1] <= sum {
		dp[n][sum] = isSubsetSum(nums, n-1, sum-nums[n-1], dp) + isSubsetSum(nums, n-1, sum, dp)
		return dp[n][sum]
	} else {
		dp[n][sum] = isSubsetSum(nums, n-1, sum, dp)
		return dp[n][sum]
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
