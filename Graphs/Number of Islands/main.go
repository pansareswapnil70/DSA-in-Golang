package main

import "fmt"

func main() {
	fmt.Println("Calculate the Number of Islands:")
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	numOfIslands := numOfIslands(grid)
	fmt.Println("Number of islands are: ", numOfIslands)
}

func numOfIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])
	answer := 0
	for i := 0; i < n; i++ {
		if len(grid[i]) != m {
			fmt.Println("Error: Irregular grid detected.")
			return -1 // Indicates an error
		}
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				answer += 1
				recursiveNumOfIslands(i, j, n, m, grid)
			}
		}
	}
	return answer
}

func recursiveNumOfIslands(i int, j int, n int, m int, grid [][]byte) {
	grid[i][j] = '0'
	if isValid(i+1, j, n, m, grid) {
		recursiveNumOfIslands(i+1, j, n, m, grid)
	}
	if isValid(i-1, j, n, m, grid) {
		recursiveNumOfIslands(i-1, j, n, m, grid)
	}
	if isValid(i, j+1, n, m, grid) {
		recursiveNumOfIslands(i, j+1, n, m, grid)
	}
	if isValid(i, j-1, n, m, grid) {
		recursiveNumOfIslands(i, j-1, n, m, grid)
	}
}

func isValid(i int, j int, n int, m int, grid [][]byte) bool {
	if i >= 0 && i < n && j >= 0 && j < m && grid[i][j] == '1' {
		return true
	}
	return false
}
