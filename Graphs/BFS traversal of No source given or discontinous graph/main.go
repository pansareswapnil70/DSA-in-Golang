package main

import "fmt"

func main() {
	fmt.Println("BFS Traversal of Disconnected graph")
	vertices := 7
	var adjacencyList = make([][]int, vertices)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 0, 2)
	addEdge(adjacencyList, 1, 3)
	addEdge(adjacencyList, 2, 3)
	addEdge(adjacencyList, 4, 5)
	addEdge(adjacencyList, 4, 6)
	addEdge(adjacencyList, 5, 6)
	printGraph(adjacencyList, vertices)
	BFSofDiscontinuedGraph(adjacencyList, vertices)
}

func BFSofDiscontinuedGraph(adjacencyList [][]int, vertices int) {
	fmt.Print("\nBFS of discontinued graph is: ")
	var visited = make([]bool, vertices)
	noOfDiscontinuedGraphs := 0
	for i := 0; i < vertices; i++ {
		if visited[i] == false {
			BFS(adjacencyList, i, &visited)
			noOfDiscontinuedGraphs += 1
		}
	}
	fmt.Println("\nNo of Discontinued Graphs are: ", noOfDiscontinuedGraphs)
}

func BFS(adjacencyList [][]int, start int, visited *[]bool) {
	queue := []int{}
	(*visited)[start] = true
	queue = append(queue, start)
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Printf(" %d ", front)
		for _, vertex := range adjacencyList[front] {
			if (*visited)[vertex] == false {
				queue = append(queue, vertex)
				(*visited)[vertex] = true
			}
		}
	}

}

func printGraph(adjacencyList [][]int, vertices int) {
	fmt.Println("The adjacency list for given graph is:")
	for i := 0; i < vertices; i++ {
		fmt.Printf("\n For vertex %d: ", i)
		for _, vertex := range adjacencyList[i] {
			fmt.Printf("%d -> ", vertex)
		}
	}
}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
	adjacencyList[v] = append(adjacencyList[v], u)
}
