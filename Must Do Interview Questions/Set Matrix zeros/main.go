package main

import "fmt"

func main() {
	fmt.Println("Set Matrix Zeros:")
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}
	fmt.Println("Input Matrix:")
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d  ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	setMatrixZeros(&matrix)
	fmt.Println("Output Matrix:")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d  ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func setMatrixZeros(matrix *[3][3]int) {
	rows := []int{}
	columns := []int{}
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				columns = append(columns, j)
			}
		}
	}

	//Update rows to 0
	for i := 0; i < len(rows); i++ {
		rowIndex := rows[i]
		for j := 0; j < m; j++ {
			matrix[rowIndex][j] = 0
		}
	}

	//Update columns to 0
	for i := 0; i < len(columns); i++ {
		columnIndex := columns[i]
		for j := 0; j < n; j++ {
			matrix[j][columnIndex] = 0
		}
	}
}
