package main

import "fmt"

func main() {
	fmt.Println("Kahn's BFS Based Algorithm:")
	vertices := 5
	var adjacencyList = make([][]int, vertices)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 0, 2)
	addEdge(adjacencyList, 1, 3)
	addEdge(adjacencyList, 2, 3)
	addEdge(adjacencyList, 1, 4)
	addEdge(adjacencyList, 3, 4)
	printGraph(adjacencyList, vertices)
	result := kahnTopologicalSort(adjacencyList, vertices)
	fmt.Println("Output of Kahn's Algorithm for given directed graph is: ", result)
}

func kahnTopologicalSort(adjacencyList [][]int, vertices int) []int {
	inDegree := make([]int, vertices)
	for i := 0; i < vertices; i++ {
		for _, neighbour := range adjacencyList[i] {
			inDegree[neighbour]++
		}
	}

	queue := []int{}
	for i := 0; i < vertices; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	var result []int
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		result = append(result, front)
		for _, vertex := range adjacencyList[front] {
			inDegree[vertex]--
			if inDegree[vertex] == 0 {
				queue = append(queue, vertex)
			}
		}
	}
	return result
}

func printGraph(adjacencyList [][]int, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Printf("For vertex %d:", i)
		for _, vertex := range adjacencyList[i] {
			fmt.Printf("  %d  ", vertex)
		}
		fmt.Print("\n")
	}
}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
}
