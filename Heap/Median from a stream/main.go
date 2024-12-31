package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []int
type minHeap []int

type medianFinder struct {
	maxHeap *maxHeap
	minHeap *minHeap
}

func main() {
	fmt.Println("Median finder from a stream")
	arr := []int{10, 20, 30, 40, 50}
	mf := constructor()
	for _, num := range arr {
		mf.Add(num)
		fmt.Printf("Added %d, Current Median: %.1f\n", num, mf.FindMedian())
	}
}

func (mf *medianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	}
	return (float64((*mf.maxHeap)[0]) + float64((*mf.minHeap)[0])) / 2.0

}

func (mf *medianFinder) Add(num int) {
	if mf.maxHeap.Len() == 0 || num <= (*mf.maxHeap)[0] {
		heap.Push(mf.maxHeap, num)
	} else {
		heap.Push(mf.minHeap, num)
	}

	if mf.maxHeap.Len() > mf.minHeap.Len()+1 {
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	} else if mf.maxHeap.Len() < mf.minHeap.Len() {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

func constructor() medianFinder {
	maxHeap := &maxHeap{}
	minHeap := &minHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return medianFinder{maxHeap: maxHeap, minHeap: minHeap}
}

func (maxHeap maxHeap) Len() int {
	return len(maxHeap)
}

func (maxHeap maxHeap) Less(i, j int) bool {
	return maxHeap[i] > maxHeap[j]
}

func (maxHeap maxHeap) Swap(i, j int) {
	maxHeap[i], maxHeap[j] = maxHeap[j], maxHeap[i]
}

func (maxHeap *maxHeap) Push(x interface{}) {
	*maxHeap = append(*maxHeap, x.(int))
}

func (maxHeap *maxHeap) Pop() interface{} {
	n := (*maxHeap).Len()
	item := (*maxHeap)[n-1]
	(*maxHeap) = (*maxHeap)[:n-1]
	return item
}

func (minHeap minHeap) Len() int {
	return len(minHeap)
}

func (minHeap minHeap) Less(i, j int) bool {
	return minHeap[i] < minHeap[j]
}

func (minHeap minHeap) Swap(i, j int) {
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *minHeap) Push(x interface{}) {
	*minHeap = append(*minHeap, x.(int))
}

func (minHeap *minHeap) Pop() interface{} {
	n := (*minHeap).Len()
	item := (*minHeap)[n-1]
	*minHeap = (*minHeap)[:n-1]
	return item
}
