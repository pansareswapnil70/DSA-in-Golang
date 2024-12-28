package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func main() {
	fmt.Println("Minimum Cost of connecting ropes:")
	ropes := []int{4, 3, 2, 6}
	minCost := minCostToConnectRopes(ropes)
	fmt.Println("Minimum cost to connect ropes:", minCost)
}

func minCostToConnectRopes(ropes []int) int {
	MinHeap := &MinHeap{}
	heap.Init(MinHeap)
	for _, rope := range ropes {
		heap.Push(MinHeap, rope)
	}
	cost := 0
	for MinHeap.Len() > 1 {
		first := heap.Pop(MinHeap).(int)
		second := heap.Pop(MinHeap).(int)

		cost += first + second

		heap.Push(MinHeap, first+second)
	}
	return cost
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := h.Len()
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}
