package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, weight int
}

type Graph struct {
	adjList map[int][]Edge
}

type Item struct {
	node, distance int
}

type priorityQueue []Item

func main() {
	fmt.Println("Djikstra's algorithm with minHeap")
	g := newGraph()
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 8)
	g.AddEdge(2, 4, 2)
	g.AddEdge(3, 4, 7)
	start := 0
	distances := g.djikstraAlgo(start)
	fmt.Println("From node: ", start)

	for node, distance := range distances {
		fmt.Printf("\nTo node %d: %d", node, distance)
	}
}

func (g Graph) djikstraAlgo(start int) map[int]int {
	dist := make(map[int]int)
	for node := range g.adjList {
		dist[node] = math.MaxInt64
	}
	visited := make(map[int]bool)
	pq := &priorityQueue{}
	heap.Init(pq)
	dist[start] = 0
	heap.Push(pq, Item{node: start, distance: 0})
	for pq.Len() > 0 {
		poppedElement := heap.Pop(pq).(Item)
		node := poppedElement.node

		if visited[node] == true {
			continue
		}
		visited[node] = true
		for _, edge := range g.adjList[node] {
			to, weight := edge.to, edge.weight
			if dist[node]+weight < dist[to] {
				dist[to] = dist[node] + weight
				heap.Push(pq, Item{node: to, distance: dist[to]})
			}
		}
	}
	return dist
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *priorityQueue) Pop() interface{} {
	n := (*pq).Len()
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func newGraph() Graph {
	return Graph{
		adjList: make(map[int][]Edge),
	}
}

func (g Graph) AddEdge(from int, to int, weight int) {
	g.adjList[from] = append(g.adjList[from], Edge{to: to, weight: weight})
	g.adjList[to] = append(g.adjList[to], Edge{from, weight})
}
