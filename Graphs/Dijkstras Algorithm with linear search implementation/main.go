package main

import (
	"fmt"
	"math"
)

func main() {
	graph := [][]int{
		{0, 5, 0, 9, 0},
		{5, 0, 2, 0, 0},
		{0, 2, 0, 3, 7},
		{9, 0, 3, 0, 0},
		{0, 0, 7, 0, 0},
	}
	fmt.Println("The given nodes are:", graph)
	start := 0
	end := 4

	dist := dijkstra(graph, start)

	fmt.Printf("Shortest path from node %d to %d: %d\n", start, end, dist[end])
}

func dijkstra(graph [][]int, start int) []int {
	n := len(graph)
	distances := make([]int, n)
	for i := 0; i < n; i++ {
		distances[i] = math.MaxInt32
	}
	visited := make([]bool, n)
	distances[start] = 0
	for i := 0; i < n; i++ {
		u := -1
		for j := 0; j < n; j++ {
			if visited[j] == false && (u == -1 || distances[j] < distances[u]) {
				u = j
			}
		}
		if u == -1 {
			break
		}
		visited[u] = true
		for v := 0; v < n; v++ {
			if graph[u][v] != 0 && visited[v] != true && distances[u]+graph[u][v] < distances[v] {
				distances[v] = distances[u] + graph[u][v]
			}
		}
	}
	return distances
}
