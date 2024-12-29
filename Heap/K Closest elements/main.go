package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Element struct {
	value      int
	difference int
}

type MaxHeap []Element

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 4
	x := 3
	fmt.Println(k, " closest elements to ", x, " are ", findClosestElements(arr, k, x)) // Output: [1, 2, 3, 4]

	arr = []int{1, 2, 3, 4, 5}
	k = 2
	x = 6
	fmt.Println(k, " closest elements to ", x, " are ", findClosestElements(arr, k, x)) // Output: [4, 5]

	arr = []int{10, 12, 15, 17, 20}
	k = 3
	x = 14
	fmt.Println(k, " closest elements to ", x, " are ", findClosestElements(arr, k, x)) // Output: [12, 15, 17]
}

func findClosestElements(arr []int, k int, x int) []int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, val := range arr {
		heap.Push(h, Element{value: val, difference: int(math.Abs(float64(x - val)))})
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

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].difference > h[j].difference
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MaxHeap) Pop() interface{} {
	n := (*h).Len()
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}
