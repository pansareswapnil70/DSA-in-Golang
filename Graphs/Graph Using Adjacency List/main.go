package main

import "fmt"

func main() {
	fmt.Println("Graph Implementation using Adjacency List:")
	vertices := 4
	var adjacencyList = make([][]int, vertices)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 0, 2)
	addEdge(adjacencyList, 1, 2)
	addEdge(adjacencyList, 2, 3)
	printGraph(adjacencyList, vertices)
}

func printGraph(adjacencyList [][]int, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Println("From vertex: ", i)
		for _, v := range adjacencyList[i] {
			fmt.Printf("%d-> ", v)
		}
		fmt.Printf("\n")
	}
}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
	adjacencyList[v] = append(adjacencyList[v], u)
}
