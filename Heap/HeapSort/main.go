package main

import "fmt"

type maxHeap struct {
	data []int
}

func main() {
	fmt.Println("Heap Sort")
	arr := []int{3, 1, 6, 5, 2, 4}
	maxHeap := &maxHeap{}
	maxHeap.heapSort(arr)
	fmt.Println("Sorted data is:")
	maxHeap.Print()
}

func (h *maxHeap) heapSort(arr []int) {
	h.buildHeap(arr)
	for i := len(h.data) - 1; i >= 1; i-- {
		h.data[i], h.data[0] = h.data[0], h.data[i]
		h.maxHeapify(0, i)
	}
}

func (h *maxHeap) buildHeap(arr []int) {
	h.data = append([]int(nil), arr...)
	for i := (len(h.data) - 2) / 2; i >= 0; i-- {
		h.maxHeapify(i, len(h.data))
	}
}

func (h *maxHeap) maxHeapify(index int, size int) {
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	largest := index
	if leftIndex < size && h.data[leftIndex] > h.data[largest] {
		largest = leftIndex
	}
	if rightIndex < size && h.data[rightIndex] > h.data[largest] {
		largest = rightIndex
	}
	if largest != index {
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		h.maxHeapify(largest, size)
	}
}

func (h *maxHeap) Print() {
	fmt.Println("Max Heap Data is:", h.data)
}
