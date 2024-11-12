package main

import (
	"fmt"
)

func main() {
	fmt.Println("All Sorting Algorithms")
	a := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted Array:", a)
	arr := bubbleSort(a)
	fmt.Println("Bubble Sorted Array:")
	fmt.Println(arr)
	b := []int{64, 34, 25, 12, 22, 11, 90}
	selection := selectionSort(b)
	fmt.Println("Selection Sorted Array:")
	fmt.Println(selection)
	c := []int{64, 34, 25, 12, 22, 11, 90}
	insertion := insertionSort(c)
	fmt.Println("Insertion Sorted Array:")
	fmt.Println(insertion)
	d := []int{64, 34, 25, 12, 22, 11, 90}
	merge := mergeSort(d)
	fmt.Println("Merge Sorted Array:")
	fmt.Println(merge)
	e := []int{64, 34, 25, 12, 22, 11, 90}
	quick := quickSort(e, 0, len(e)-1)
	fmt.Println("Quick Sorted Array:")
	fmt.Println(quick)
}

func quickSort(a []int, start int, end int) []int {
	if start >= end {
		return a
	}
	pivot := partition(a, start, end)
	quickSort(a, start, pivot-1)
	quickSort(a, pivot+1, end)
	return a
}

func partition(a []int, start int, end int) int {
	pivot := a[end] // Choose the last element as the pivot
	i := start - 1  // This will be the index of the smaller element

	for j := start; j < end; j++ {
		if a[j] < pivot { // If the current element is smaller than the pivot
			i++
			a[i], a[j] = a[j], a[i] // Swap the elements
		}
	}
	a[i+1], a[end] = a[end], a[i+1]

	return i + 1
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	mid := len(a) / 2
	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])

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
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}

func insertionSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
	return a
}

func selectionSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	return a
}

func bubbleSort(a []int) []int {
	n := len(a)
	end := n - 1
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < end; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}
		end = end - 1
		if !swapped {
			break
		}
	}
	return a
}
