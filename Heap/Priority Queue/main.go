package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string
	priority int
}

type priorityQueue []*Item

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	Item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return Item

}

func main() {
	fmt.Println("Priority Queue using heap") // Priority queue uses minheap in background
	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: "Task 1", priority: 4})
	heap.Push(pq, &Item{value: "Task 2", priority: 3})
	heap.Push(pq, &Item{value: "Task 3", priority: 1})
	heap.Push(pq, &Item{value: "Task 4", priority: 2})

	fmt.Printf("Top priority item: %s\n", (*pq)[0].value)

	fmt.Println("Processing tasks:")
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Printf("Task: %s (Priority: %d)\n", item.value, item.priority)
	}

}
