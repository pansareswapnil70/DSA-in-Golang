package main

import (
	"fmt"
)

func main() {
	fmt.Println("Insertion Sort")
	fmt.Println("Unsorted Array:")
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println("Sorted Array")
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	n := len(arr) // Get the length of the array

	// Outer loop: iterate over each element starting from the second element (i = 1)
	for i := 1; i < n; i++ {
		key := arr[i] // Store the current element as `key`
		j := i - 1    // Start with the element to the left of `key`

		// Inner loop: shift elements greater than `key` to the right
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // Shift element one position to the right
			j -= 1            // Move left in the array
		}

		// Insert `key` in its correct position in the sorted portion
		arr[j+1] = key
	}
}
