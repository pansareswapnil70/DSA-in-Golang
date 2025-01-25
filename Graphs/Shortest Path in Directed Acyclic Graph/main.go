package main

import (
	"fmt"
	"math"
)

func main() {
	// Number of vertices
	vertices := 6

	// Graph represented as an adjacency list with weights
	graph := make([][]Edge, vertices)
	addEdge(graph, 0, 1, 2)
	addEdge(graph, 0, 4, 1)
	addEdge(graph, 1, 2, 3)
	addEdge(graph, 4, 2, 2)
	addEdge(graph, 4, 5, 4)
	addEdge(graph, 5, 3, 1)
	addEdge(graph, 2, 3, 6)

	// Print graph
	fmt.Println("Graph adjacency list with weights:")
	printGraph(graph, vertices)

	// Compute shortest paths from vertex 0
	start := 0
	distances := shortestPathDAG(graph, vertices, start)

	// Print shortest paths
	fmt.Printf("\nShortest paths from vertex %d:\n", start)
	for i, d := range distances {
		if d == math.MaxInt {
			fmt.Printf("Vertex %d: Unreachable\n", i)
		} else {
			fmt.Printf("Vertex %d: %d\n", i, d)
		}
	}
}

// Edge represents a directed edge with a weight
type Edge struct {
	to     int
	weight int
}

// addEdge adds a directed edge with weight to the graph
func addEdge(graph [][]Edge, from, to, weight int) {
	graph[from] = append(graph[from], Edge{to: to, weight: weight})
}

// topologicalSort computes a topological order of the graph using DFS
func topologicalSort(graph [][]Edge, vertices int) []int {
	visited := make([]bool, vertices)
	stack := []int{}

	var dfs func(v int)
	dfs = func(v int) {
		visited[v] = true
		for _, edge := range graph[v] {
			if !visited[edge.to] {
				dfs(edge.to)
			}
		}
		stack = append([]int{v}, stack...) // Push to the front
	}

	for i := 0; i < vertices; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	return stack
}

// shortestPathDAG finds the shortest path from a start vertex in a DAG
func shortestPathDAG(graph [][]Edge, vertices, start int) []int {
	// Initialize distances to infinity
	distances := make([]int, vertices)
	for i := range distances {
		distances[i] = math.MaxInt
	}
	distances[start] = 0

	// Get the topological order of the graph
	topOrder := topologicalSort(graph, vertices)

	// Relax edges in topological order
	for _, u := range topOrder {
		if distances[u] != math.MaxInt {
			for _, edge := range graph[u] {
				if distances[u]+edge.weight < distances[edge.to] {
					distances[edge.to] = distances[u] + edge.weight
				}
			}
		}
	}

	return distances
}

// printGraph prints the adjacency list representation of the graph
func printGraph(graph [][]Edge, vertices int) {
	for i := 0; i < vertices; i++ {
		fmt.Printf("Vertex %d: ", i)
		for _, edge := range graph[i] {
			fmt.Printf("-> (%d, weight %d) ", edge.to, edge.weight)
		}
		fmt.Println()
	}
}
