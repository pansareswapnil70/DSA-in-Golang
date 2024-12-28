package main

import (
	"container/heap"
	"fmt"
)

type priorityQueue []int

func main() {
	fmt.Println("Kth Largest Element:")
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 3
	element := KthLargestElement(nums, k)
	fmt.Println(element)
}

func KthLargestElement(nums []int, k int) int {
	pq := &priorityQueue{}
	heap.Init(pq)
	for _, num := range nums {
		heap.Push(pq, num)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	item := heap.Pop(pq).(int)
	return item
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *priorityQueue) Pop() interface{} {
	n := (*pq).Len()
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item

}
