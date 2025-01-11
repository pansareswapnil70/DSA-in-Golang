package main

import "fmt"

func main() {
	fmt.Println("Determine if path exists in graph:")
	vertices := 6
	var adjacencylist = make([][]int, vertices)
	addEdge(adjacencylist, 0, 1)
	addEdge(adjacencylist, 0, 2)
	addEdge(adjacencylist, 1, 3)
	addEdge(adjacencylist, 2, 3)
	addEdge(adjacencylist, 4, 5)
	boolResult := isPathExists(adjacencylist, 0, 3)
	fmt.Println("Does path exists between 0 and 3 ? ", boolResult)
	boolResult = isPathExists(adjacencylist, 0, 5)
	fmt.Println("Does path exists between 0 and 5 ? ", boolResult)
}

func addEdge(adjacencylist [][]int, u int, v int) {
	adjacencylist[u] = append(adjacencylist[u], v)
	adjacencylist[v] = append(adjacencylist[v], u)
}

func isPathExists(adjacencylist [][]int, start int, end int) bool {
	if start == end {
		return true
	}
	v := len(adjacencylist)
	var visited = make(map[int]bool, v+1)
	queue := []int{}
	queue = append(queue, start)
	visited[start] = true
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, vertex := range adjacencylist[front] {
			if vertex == end {
				return true
			}
			if visited[vertex] == false {
				queue = append(queue, vertex)
				visited[vertex] = true
			}
		}
	}
	return false
}
