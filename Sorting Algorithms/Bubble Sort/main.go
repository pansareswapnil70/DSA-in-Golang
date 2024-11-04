package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bubble Sort Algorithm")
	fmt.Println("Unsorted Array:")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(arr)
	sortedArr := bubbleSort(arr)
	fmt.Println("Sorted Array:")
	fmt.Println(sortedArr)
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	end := n - 1

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < end; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		end -= 1
		if !swapped {
			break
		}
	}
	return arr
}
