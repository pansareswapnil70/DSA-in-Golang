package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	value     int
	frequency int
}

type MinHeap []Element

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 3, 3, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Println(result)
}

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	h := &MinHeap{}
	heap.Init(h)
	for val, freq := range frequencyMap {
		heap.Push(h, Element{value: val, frequency: freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, 0, k)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(Element).value)
	}
	return result
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].frequency < h[j].frequency
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
