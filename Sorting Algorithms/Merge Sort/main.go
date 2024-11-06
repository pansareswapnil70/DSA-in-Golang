package main

import (
	"fmt"
)

func main() {
	fmt.Println("Merge Sort")
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted Array:")
	fmt.Println(arr)
	result := mergeSort(arr)
	fmt.Println("Sorted Array:")
	fmt.Println(result)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
	}

	for j < len(right) {
		result = append(result, right[j])
	}

	return result
}
