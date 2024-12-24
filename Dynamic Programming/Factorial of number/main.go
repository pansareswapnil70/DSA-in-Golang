package main

import (
	"fmt"
)

//Input: n = 5
// Output: 120
// Explanation: 5*4*3*2*1 = 120

// Input: n = 4
// Output: 24
// Explanation: 4*3*2*1 = 24

func main() {
	fmt.Println("Factorial of a number:")
	n := 5
	result := Factorial(n)
	fmt.Println("Factorial of ", n, " is ", result)
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}
