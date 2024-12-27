package main

import (
	"fmt"
)

type minHeap struct {
	data []int
}

func main() {
	fmt.Println("Build min heap from an array")
	minHeap := &minHeap{}
	arr := []int{3, 1, 6, 5, 2, 4}
	minHeap.buildHeap(arr)
	fmt.Println("Build Heap for array:", arr)
	minHeap.Print()
}

func (h *minHeap) buildHeap(arr []int) {
	h.data = append([]int(nil), arr...)
	for i := (len(h.data) - 2) / 2; i >= 0; i-- {
		h.minHeapify(i)
	}
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

	if smallest != index {
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		h.minHeapify(smallest)
	}
}

func (h *minHeap) Print() {
	fmt.Println("Heap data:", h.data)
}
