package main

import (
	"fmt"
	"math"
)

type minHeap struct {
	data []int
}

func main() {
	fmt.Println("Min Heap Implementation")
	minHeap := &minHeap{}
	minHeap.insert(3)
	minHeap.insert(1)
	minHeap.insert(6)
	minHeap.insert(5)
	minHeap.insert(2)
	minHeap.insert(4)
	fmt.Println("Min Heap data before heapify:")
	minHeap.print()
	minHeap.minHeapify(0)
	fmt.Println("Min Heap data after heapify:")
	minHeap.print()
	min := minHeap.getMin()
	fmt.Println("Minimum value in the heap is: ", min)
	extractedMin := minHeap.extractMin()
	fmt.Println("Extracted value from the heap is ", extractedMin, " and heap size is now: ", len(minHeap.data))
	minHeap.print()
	min = minHeap.getMin()
	fmt.Println("Minimum value in the heap is:", min)
}

func (h *minHeap) extractMin() int {
	if len(h.data) == 0 {
		return math.MaxInt
	}
	if len(h.data) == 1 {
		min := h.data[0]
		h.data = h.data[:len(h.data)-1]
		return min
	}
	min := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1] // Remove the last element
	h.minHeapify(0)                 // Restore the heap property
	return min
}

func (h *minHeap) getMin() int {
	return h.data[0]
}

func (h *minHeap) minHeapify(index int) {
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	smallest := index

	if leftIndex < len(h.data) && h.data[leftIndex] < h.data[smallest] {
		smallest = leftIndex
	}

	if rightIndex < len(h.data) && h.data[rightIndex] < h.data[smallest] {
		smallest = rightIndex
	}

	if index != smallest {
		h.data[smallest], h.data[index] = h.data[index], h.data[smallest]
		h.minHeapify(smallest)
	}
}

func (h *minHeap) insert(index int) {
	h.data = append(h.data, index)
	h.bubbleUp(len(h.data) - 1)
}

func (h *minHeap) bubbleUp(index int) {
	parent := (index - 1) / 2
	for index > 0 && h.data[index] < h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
		parent = (index - 1) / 2
	}
}

func (h *minHeap) print() {
	fmt.Println(h.data)
}
