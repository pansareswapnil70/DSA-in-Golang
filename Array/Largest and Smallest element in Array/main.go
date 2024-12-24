package main

import (
	"fmt"
)

func main() {
	fmt.Println("Largest and Smallest Element in an array:")
	nums := []int{12, 5, 3, 7, 19, 1, 8, 10}
	smallest, largest := largestAndSmallest(nums)
	fmt.Println("Largest and smallest elements are ", largest, " and ", smallest)
}

func largestAndSmallest(nums []int) (int, int) {
	n := len(nums)
	if n == 0 {
		return 0, 0
	}

	min, max := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return min, max
}
