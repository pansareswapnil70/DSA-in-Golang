package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array Subset Sum:")
	arr := []int{3, 34, 4, 12, 5, 2}
	sum := 9
	n := len(arr)
	boolResult := isSubsetSum(arr, sum, n)
	fmt.Println("Subset sum is", boolResult)
}

func isSubsetSum(arr []int, sum int, n int) bool {
	if n == 0 {
		return false
	}
	if sum == 0 {
		return true
	}

	if arr[n-1] <= sum {
		return isSubsetSum(arr, sum-arr[n-1], n-1) || isSubsetSum(arr, sum, n-1)
	} else {
		return isSubsetSum(arr, sum, n-1)
	}
}
