package main

import (
	"fmt"
)

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Quick Sort")
	fmt.Println("Unsorted Array:")
	fmt.Println(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted Array:")
	fmt.Println(arr)
}

func quickSort(arr []int, start int, end int) {
	if end <= start {
		return
	}
	if start < end {
		pivot := partition(arr, start, end)
		quickSort(arr, start, pivot-1)
		quickSort(arr, pivot+1, end)
	}
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
