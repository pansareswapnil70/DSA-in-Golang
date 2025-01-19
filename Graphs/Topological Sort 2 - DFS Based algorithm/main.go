package main

import "fmt"

func main() {
	fmt.Println("Topological Sort 2 - DFS based Algorithm")
	vertices := 5
	adjacencyList := make([][]int, vertices)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 1, 3)
	addEdge(adjacencyList, 2, 3)
	addEdge(adjacencyList, 2, 4)
	addEdge(adjacencyList, 3, 4)
	printGraph(adjacencyList, vertices)
	topologicalSort(adjacencyList, vertices)
}

func topologicalSort(adjacencyList [][]int, vertices int) {
	visited := make([]bool, vertices)
	stack := []int{}
	for i := 0; i < vertices; i++ {
		if visited[i] == false {
			DFSRec(i, visited, &stack, adjacencyList)
		}
	}
	fmt.Println("\nDFS based topological Sort:")
	for len(stack) > 0 {
		fmt.Printf(" %d ", stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
}

func DFSRec(start int, visited []bool, stack *[]int, adjacencyList [][]int) {
	visited[start] = true
	for _, neighbour := range adjacencyList[start] {
		if visited[neighbour] == false {
			DFSRec(neighbour, visited, stack, adjacencyList)
		}
	}
	*stack = append(*stack, start)
}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
}

func printGraph(adjacencyList [][]int, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Printf("For vertex %d: ", i)
		for _, neighbour := range adjacencyList[i] {
			fmt.Printf(" %d ", neighbour)
		}
		fmt.Printf("\n")
	}
}
