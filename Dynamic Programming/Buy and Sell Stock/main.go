package main

import (
	"fmt"
)

func main() {
	fmt.Println("Best Time to Buy and Sell Stock")
	arr := []int{7, 1, 5, 3, 6, 4}
	maxProfit := calculateMaxProfit(arr)
	fmt.Println("Max profit obtained is: ", maxProfit)
}

func calculateMaxProfit(arr []int) int {
	maxProfit := 0
	buy := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < buy {
			buy = arr[i]
		} else {
			currentProfit := arr[i] - buy
			maxProfit = max(maxProfit, currentProfit)
		}
	}
	return maxProfit
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
