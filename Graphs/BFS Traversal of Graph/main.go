package main

import "fmt"

func main() {
	fmt.Println("BFS Traversal of Graph:")
	vertices := 5
	var adjacencyList = make([][]int, vertices)
	addEdge(adjacencyList, 0, 1)
	addEdge(adjacencyList, 0, 2)
	addEdge(adjacencyList, 1, 2)
	addEdge(adjacencyList, 1, 3)
	addEdge(adjacencyList, 2, 3)
	addEdge(adjacencyList, 2, 4)
	addEdge(adjacencyList, 3, 4)
	printGraph(adjacencyList, vertices)
	fmt.Println("BFS Traversal of Graph is: ")
	source := 0
	BFSTraversal(adjacencyList, vertices, source)
}

func BFSTraversal(adjacencyList [][]int, v int, s int) {
	var visited = make([]bool, v+1)
	queue := []int{}
	queue = append(queue, s)
	visited[s] = true
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Printf(" %d ", front)
		for _, vertex := range adjacencyList[front] {
			if visited[vertex] == false {
				queue = append(queue, vertex)
				visited[vertex] = true
			}
		}
	}
}

func addEdge(adjacencyList [][]int, u int, v int) {
	adjacencyList[u] = append(adjacencyList[u], v)
	adjacencyList[v] = append(adjacencyList[v], u)
}

func printGraph(adjacencyList [][]int, v int) {
	for i := 0; i < v; i++ {
		fmt.Printf("For vertex:%d ", i)
		for _, vertex := range adjacencyList[i] {
			fmt.Printf(" %d->", vertex)
		}
		fmt.Printf("\n")
	}
}
