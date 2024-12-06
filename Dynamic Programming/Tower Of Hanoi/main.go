package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tower Of Hanoi:")
	noOfDisks := []int{1, 2, 3, 4, 5, 6}
	var from, aux, to = 1, 2, 3
	for _, n := range noOfDisks {
		moves := 0
		towerOfHanoi(n, from, to, aux, &moves)
		fmt.Println("No of moves required:", moves, " for ", n, "disks.")
	}
}

// Function to solve the Tower of Hanoi problem recursively
// Parameters:
// n     : number of disks to move
// from  : the rod from which disks are to be moved
// to    : the target rod to which disks are to be moved
// aux   : the auxiliary rod used for intermediate steps
// moves : pointer to the integer variable tracking the number of moves
func towerOfHanoi(n int, from int, to int, aux int, moves *int) {
	// Base case: If there is only one disk, just move it
	if n <= 1 {
		// Increment the move count
		(*moves)++
		return
	}
	// Step 1: Move n-1 disks from the source rod to the auxiliary rod
	towerOfHanoi(n-1, from, aux, to, moves)
	// Step 2: Move the nth (largest) disk from the source rod to the target rod
	(*moves)++ // Increment the move count for this disk movement
	// Step 3: Move the n-1 disks from the auxiliary rod to the target rod
	towerOfHanoi(n-1, aux, to, from, moves)
}
