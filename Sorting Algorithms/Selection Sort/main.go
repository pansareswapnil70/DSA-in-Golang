package main

import (
	"fmt"
)

func main() {
	fmt.Println("Selection Sort")
	fmt.Println("Unsorted Array:")
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println(arr)
	selectionSort(arr)
	fmt.Println("Sorted Array:")
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
