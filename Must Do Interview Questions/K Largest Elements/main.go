package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type minHeap []int

func main() {
	fmt.Println("K Largest Element from an array")
	arr := []int{3, 4, 2, 1}
	k := 2
	n := 4
	fmt.Println("1st Approach using sort package:")
	sort.Ints(arr)
	fmt.Println("K Largest Elements are: ", arr[n-k:n])

	fmt.Println("2nd Approach using heap:")
	h := &minHeap{}
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := []int{}
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(int))
	}
	fmt.Println("K Largest Elements using Heap are: ", result)
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}
