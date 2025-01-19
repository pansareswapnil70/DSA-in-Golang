package main

import "fmt"

func main() {
	fmt.Println("Detect cycle in directed graph:")
	vertices := 8
	adjacencyList := make([][]int, vertices)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 1, 2)
	addEdge(adjacencyList, 2, 3)
	addEdge(adjacencyList, 1, 4)
	addEdge(adjacencyList, 4, 3)
	addEdge(adjacencyList, 5, 0)
	addEdge(adjacencyList, 5, 6)
	addEdge(adjacencyList, 6, 7)
	addEdge(adjacencyList, 7, 5)
	printGraph(adjacencyList, vertices)
	boolResult := boolIsCyclic(adjacencyList, vertices)
	if boolResult {
		fmt.Println("Given graph contains cycle")
	} else {
		fmt.Println("Given graph does not contain cycle")
	}
}

func cycle(start int, adjacencyList [][]int, visited []bool, dfsVisited []bool, vertices int) bool {
	visited[start] = true
	dfsVisited[start] = true
	for _, neighbour := range adjacencyList[start] {
		if visited[neighbour] == false {
			if cycle(neighbour, adjacencyList, visited, dfsVisited, vertices) {
				return true
			}
		} else if visited[neighbour] == true && dfsVisited[neighbour] == true {
			return true
		}
	}
	dfsVisited[start] = false
	return false
}

func boolIsCyclic(adjacencyList [][]int, vertices int) bool {
	visited := make([]bool, vertices)
	dfsVisited := make([]bool, vertices)
	for i := 0; i < vertices; i++ {
		if visited[i] == false {
			if cycle(i, adjacencyList, visited, dfsVisited, vertices) {
				return true
			}
		}
	}
	return false
}

func printGraph(adjacencyList [][]int, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Printf("For vertex %d: ", i)
		for _, vertex := range adjacencyList[i] {
			fmt.Printf(" %d ", vertex)
		}
		fmt.Printf("\n")
	}
}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
}
