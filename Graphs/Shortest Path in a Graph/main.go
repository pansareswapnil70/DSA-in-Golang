package main

import "fmt"

func main() {
	fmt.Println("Shortest Path in a Graph:")
	vertices := 4
	adjList := make([][]int, vertices)
	addEdge(adjList, 0, 1)
	addEdge(adjList, 0, 3)
	addEdge(adjList, 1, 2)
	addEdge(adjList, 1, 3)
	addEdge(adjList, 2, 3)
	printGraph(adjList, vertices)
	start := 0
	result := findShortestPaths(start, adjList, vertices)
	for i, distance := range result {
		fmt.Printf("\nShortest path between %d and %d is: %d", start, i, distance)
	}
}

func findShortestPaths(start int, adjList [][]int, vertices int) []int {
	distance := make([]int, vertices)
	visited := make([]bool, vertices)
	for i := 0; i < vertices; i++ {
		distance[i] = -1 // Initialize distances as -1 (indicating unreachable)
	}
	visited[start] = true
	distance[start] = 0
	queue := []int{}
	queue = append(queue, start)
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, neighbour := range adjList[front] {
			if visited[neighbour] == false {
				distance[neighbour] = distance[front] + 1
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
	return distance
}

func printGraph(adjList [][]int, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Printf("For vertex %d: ", i)
		for _, vertex := range adjList[i] {
			fmt.Printf(" %d ", vertex)
		}
		fmt.Printf("\n")
	}
}

func addEdge(adjList [][]int, u int, v int) {
	adjList[u] = append(adjList[u], v)
	adjList[v] = append(adjList[v], u)
}
