package main

import (
	"fmt"
)

type minHeap struct {
	data []int
}

func main() {
	fmt.Println("Build heap from an array")
	arr := []int{3, 1, 6, 5, 2, 4}
	minHeap := &minHeap{}
	minHeap.buildHeap(arr)
	fmt.Println("Minheap built from array", arr, " is ", minHeap.data)
}

func (h *minHeap) buildHeap(arr []int) {
	h.data = append([]int(nil), arr...)
	for i := (len(h.data) / 2) - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *minHeap) heapify(index int) {
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	smallest := index
	if leftIndex < len(h.data) && h.data[leftIndex] < h.data[index] {
		smallest = leftIndex
	}
	if rightIndex < len(h.data) && h.data[rightIndex] < h.data[index] {
		smallest = rightIndex
	}
	if index != smallest {
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		h.heapify(smallest)
	}
}
