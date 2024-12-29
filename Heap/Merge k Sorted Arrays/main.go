package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	value      int
	arrayIndex int
	valueIndex int
}

type MinHeap []Element

func main() {
	// Test Case
	arrays := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	result := mergeKSortedArrays(arrays)
	fmt.Println(result) // Output: [1, 2, 3, 4, 5, 6, 7, 8, 9]
}

func mergeKSortedArrays(arrays [][]int) []int {
	h := &MinHeap{}
	heap.Init(h)
	for i, arr := range arrays {
		heap.Push(h, Element{value: arr[0], arrayIndex: i, valueIndex: 0})
	}
	result := []int{}
	for h.Len() > 0 {
		smallest := heap.Pop(h).(Element)
		result = append(result, smallest.value)
		if smallest.valueIndex+1 < len(arrays[smallest.arrayIndex]) {
			nextValue := arrays[smallest.arrayIndex][smallest.valueIndex+1]
			heap.Push(h, Element{value: nextValue, arrayIndex: smallest.arrayIndex, valueIndex: smallest.valueIndex + 1})
		}
	}
	return result
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	n := (*h).Len()
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}
