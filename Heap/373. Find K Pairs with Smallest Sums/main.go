package main

import (
	"container/heap"
	"fmt"
)

type ElementPairs struct {
	Pairs []int
	Sum   int
}

type minHeap []ElementPairs

func main() {
	fmt.Println("Find K Pairs with smallest sums")
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	h := &minHeap{}
	heap.Init(h)
	result := h.findKPairs(nums1, nums2, k)
	fmt.Println(k, " pairs with smallest sum are: ", result)
}

func (h *minHeap) findKPairs(nums1 []int, nums2 []int, k int) [][]int {
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			slice := []int{nums1[i], nums2[j]}
			heap.Push(h, ElementPairs{
				Pairs: slice,
				Sum:   sum,
			})
		}
	}

	var result [][]int
	for i := 0; i < k && h.Len() > 0; i++ {
		element := heap.Pop(h).(ElementPairs)
		pair := element.Pairs
		result = append(result, pair)
	}
	return result
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].Sum < h[j].Sum
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(ElementPairs))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}
