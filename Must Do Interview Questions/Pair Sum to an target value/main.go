package main

import "fmt"

func main() {
	nums := []int{2, 4, 3, 5, 7, 8, -1, 1}
	target := 6
	fmt.Println("Retuns pairs of elements from an array that sum to a target value: ", target)
	result := pairSum(nums, target)
	fmt.Println("Following pairs sum equals to target value:", result)
}

func pairSum(nums []int, target int) [][2]int {
	var pairs = [][2]int{}
	var seen = make(map[int]bool)

	for _, num := range nums {
		complement := target - num
		if seen[complement] {
			pairs = append(pairs, [2]int{num, complement})
		}
		seen[num] = true
	}

	return pairs
}
