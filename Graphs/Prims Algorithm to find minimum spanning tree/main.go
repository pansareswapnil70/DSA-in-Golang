package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	to     int
	weight int
}

type PriorityQueue []Edge

func main() {
	fmt.Println("Prim's Algorithm to find minimum spanning tree:")
	vertices := 4
	adjacencyList := make(map[int][]Edge, vertices)
	addEdge(adjacencyList, 0, 1, 4)
	addEdge(adjacencyList, 0, 2, 6)
	addEdge(adjacencyList, 1, 2, 5)
	addEdge(adjacencyList, 1, 3, 7)
	addEdge(adjacencyList, 2, 3, 8)
	printGraph(adjacencyList, vertices)
	// Call Prim's algorithm.
	result := primMST(adjacencyList, vertices)
	fmt.Printf("Total weight of the Minimum Spanning Tree: %d\n", result)
}

func primMST(adjacencyList map[int][]Edge, vertices int) int {
	visited := make([]bool, vertices)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Edge{to: 0, weight: 0})
	totalWeight := 0
	for pq.Len() > 0 {
		edge := heap.Pop(pq).(Edge)
		if visited[edge.to] == true {
			continue
		}
		visited[edge.to] = true
		totalWeight += edge.weight
		for _, neighbour := range adjacencyList[edge.to] {
			if visited[neighbour.to] == false {
				heap.Push(pq, Edge{to: neighbour.to, weight: neighbour.weight})
			}
		}

	}
	return totalWeight
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := pq.Len()
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func addEdge(adjacencyList map[int][]Edge, u int, v int, weight int) {
	adjacencyList[u] = append(adjacencyList[u], Edge{to: v, weight: weight})
	adjacencyList[v] = append(adjacencyList[v], Edge{to: u, weight: weight})
}

func printGraph(adjacencyList map[int][]Edge, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Printf("For vertex %d:", i)
		for _, vertex := range adjacencyList[i] {
			fmt.Printf("{ %d (weight : %d) }", vertex.to, vertex.weight)
		}
		fmt.Printf("\n")
	}
}
