package main

import "fmt"

func main() {
	fmt.Println("Iterative DFS of graph: ")
	vertices := 5
	var adjacencyList = make([][]int, vertices)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 0, 2)
	addEdge(adjacencyList, 1, 3)
	addEdge(adjacencyList, 2, 3)
	addEdge(adjacencyList, 1, 4)
	addEdge(adjacencyList, 3, 4)
	printGraph(adjacencyList, vertices)
	iterativeDFS(adjacencyList, 0, vertices)
}

func iterativeDFS(adjacencyList [][]int, start int, vertices int) {
	if start < 0 || start >= vertices {
		fmt.Println("Invalid start vertex")
		return
	}

	var visited = make(map[int]bool, vertices)
	var stack = []int{}
	stack = append(stack, start)
	visited[start] = true
	fmt.Println("Iterative DFS traversal: ")
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf(" %d ", top)
		for _, vertex := range adjacencyList[top] {
			if visited[vertex] == false {
				stack = append(stack, vertex)
				visited[vertex] = true
			}
		}
	}
}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
	adjacencyList[v] = append(adjacencyList[v], u)
}

func printGraph(adjacencyList [][]int, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Printf("\n For vertex %d:", i)
		for _, vertex := range adjacencyList[i] {
			fmt.Printf(" %d ->", vertex)
		}
		fmt.Printf("\n")
	}
}
