package main

import "fmt"

func main() {
	fmt.Println("Rod Cutting Problem")
	profits := []int{1, 5, 8, 9, 10, 17, 17, 20}
	n := 8
	w := 8
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			dp[i][j] = -1
		}
	}
	maxProfit := calculateMaxProfit(profits, n, w, dp)
	fmt.Println("Maximum Profit is: ", maxProfit)
}

func calculateMaxProfit(profits []int, n int, w int, dp [][]int) int {
	if n == 0 || w == 0 {
		dp[n][w] = 0
		return dp[n][w]
	}
	if dp[n][w] != -1 {
		return dp[n][w]
	}
	if n <= w {
		dp[n][w] = max(profits[n-1]+calculateMaxProfit(profits, n-1, w-n, dp), calculateMaxProfit(profits, n-1, w, dp))
	} else {
		dp[n][w] = calculateMaxProfit(profits, n-1, w, dp)
	}
	return dp[n][w]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
