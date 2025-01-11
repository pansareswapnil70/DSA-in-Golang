package main

import "fmt"

func main() {
	fmt.Println("Find the town judge:")
	n := 3
	trust := [][]int{{1, 3}, {2, 3}}
	fmt.Println(findJudge(n, trust))
}

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}

	var trustCounts = make([]int, n+1)

	for _, t := range trust {
		trustCounts[t[0]]--
		trustCounts[t[1]]++
	}

	for i := 1; i <= n; i++ {
		if trustCounts[i] == n-1 {
			return i
		}
	}

	return -1
}
