package main

import (
	"fmt"
)

type minHeap struct {
	data []int
}

func main() {
	fmt.Println("Min Heap")
	minHeap := &minHeap{
		data: []int{3, 1, 6, 5, 2, 4},
	}
	fmt.Println("Before minHeapify:")
	minHeap.Print()
	minHeap.minHeapify(0)
	fmt.Println("Before minHeapify:")
	minHeap.Print()
}

func (h *minHeap) minHeapify(index int) {
	smallest := index
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	if leftIndex < len(h.data) && h.data[leftIndex] < h.data[smallest] {
		smallest = leftIndex
	}
	if rightIndex < len(h.data) && h.data[rightIndex] < h.data[smallest] {
		smallest = rightIndex
	}

	if smallest != index {
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		h.minHeapify(smallest)
	}
}

func (h *minHeap) insert(Val int) {
	h.data = append(h.data, Val)
	h.bubbleUp(len(h.data) - 1)
}

func (h *minHeap) bubbleUp(index int) {
	parent := (index - 1) / 2
	if parent >= 0 && h.data[index] < h.data[parent] {
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		h.bubbleUp(h.data[parent])
	}
}

func (h *minHeap) Print() {
	fmt.Println("Heap elements:", h.data)
}
