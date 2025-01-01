package main

import "fmt"

type minHeap struct {
	data []int
}

func main() {
	fmt.Println("Heap Sort an array:")
	arr := []int{3, 1, 6, 5, 2, 4}
	h := &minHeap{}
	result := h.heapSort(arr)
	fmt.Println("Input array is:", arr)
	fmt.Println("Sorted array is: ", result)
}

func (h *minHeap) heapSort(arr []int) []int {
	h.buildHeap(arr)
	n := len(h.data)
	for i := n - 1; i >= 1; i-- {
		h.data[0], h.data[i] = h.data[i], h.data[0]
		h.minHeapify(0, i)
	}
	result := []int{}
	for j := n - 1; j >= 0; j-- {
		result = append(result, h.data[j])
	}
	return result
}

func (h *minHeap) buildHeap(arr []int) {
	h.data = append([]int(nil), arr...)
	for i := (len(h.data) - 2) / 2; i >= 0; i-- {
		h.minHeapify(i, len(h.data))
	}
}

func (h *minHeap) minHeapify(i int, size int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < size && h.data[left] < h.data[smallest] {
		smallest = left
	}
	if right < size && h.data[right] < h.data[smallest] {
		smallest = right
	}
	if i != smallest {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.minHeapify(smallest, size)
	}

}
