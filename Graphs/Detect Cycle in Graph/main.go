package main

import "fmt"

func main() {
	fmt.Println("Detect Cycle in Graph")
	v := 6
	var adjacencyList = make([][]int, v)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 1, 2)
	addEdge(adjacencyList, 1, 3)
	addEdge(adjacencyList, 2, 3)
	addEdge(adjacencyList, 2, 4)
	addEdge(adjacencyList, 4, 5)
	printGraph(adjacencyList, v)
	boolResult := DFS(adjacencyList, v)
	fmt.Println("\nGiven graph has cycle?", boolResult)
}

func DFS(adjacencyList [][]int, v int) bool {
	var visited = make([]bool, v)
	for i := 0; i < v; i++ {
		if visited[i] == false {
			if DFSRecursive(adjacencyList, i, visited, -1) {
				return true
			}
		}
	}
	return false
}

func DFSRecursive(adjacencyList [][]int, vertex int, visited []bool, parent int) bool {
	visited[vertex] = true
	for _, neighbour := range adjacencyList[vertex] {
		if visited[neighbour] == false {
			if DFSRecursive(adjacencyList, neighbour, visited, vertex) {
				return true
			}
		} else if vertex != parent {
			return true
		}
	}
	return false
}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
	adjacencyList[v] = append(adjacencyList[v], u)
}

func printGraph(adjacencyList [][]int, v int) {
	for i := 0; i < v; i++ {
		fmt.Printf("\n For vertex: %d vertices are: ", i)
		for _, vertex := range adjacencyList[i] {
			fmt.Printf(" %d ", vertex)
		}
	}
}
