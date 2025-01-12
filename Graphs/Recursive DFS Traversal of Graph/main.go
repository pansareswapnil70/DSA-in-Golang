package main

import "fmt"

func main() {
	fmt.Println("DFS traversal of graph:")
	vertices := 5
	var adjacencyList = make([][]int, vertices)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 0, 2)
	addEdge(adjacencyList, 1, 3)
	addEdge(adjacencyList, 2, 3)
	addEdge(adjacencyList, 1, 4)
	addEdge(adjacencyList, 3, 4)
	printGraph(adjacencyList, vertices)
	DFS(adjacencyList, vertices, 0)
}

func DFS(adjacencyList [][]int, v int, start int) {
	var visited = make(map[int]bool, v+1)
	for i := 0; i < v; i++ {
		if visited[i] == false {
			fmt.Printf("\nDFS traversal for Source %d:", i)
			DFSRecursive(adjacencyList, i, visited)
		}
	}
}

func DFSRecursive(adjacencyList [][]int, start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("%d->", start)
	for _, v := range adjacencyList[start] {
		if visited[v] == false {
			DFSRecursive(adjacencyList, v, visited)
		}
	}
}

func printGraph(adjacencyList [][]int, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Printf("For vertex %d: ", i)
		for _, vertex := range adjacencyList[i] {
			fmt.Printf(" %d-> ", vertex)
		}
		fmt.Printf("\n")
	}

}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
	adjacencyList[v] = append(adjacencyList[v], u)
}
