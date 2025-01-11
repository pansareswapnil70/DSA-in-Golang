package main

import (
	"fmt"
)

func main() {
	fmt.Println("Beautiful String")
	inputString := "0000"
	n := len(inputString)
	generatedString1 := generate(n, true)
	diff1 := change(inputString, generatedString1)
	generatedString2 := generate(n, false)
	diff2 := change(inputString, generatedString2)
	result := min(diff1, diff2)
	fmt.Println("Minimum number of changes required: ", result)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func change(inputString string, generatedString string) int {
	diff := 0
	for i := 0; i < len(inputString); i++ {
		if inputString[i] != generatedString[i] {
			diff += 1
		}
	}
	return diff
}

func generate(n int, zero bool) string {
	outputString := ""
	if n == 0 {
		return outputString
	}

	if zero {
		outputString += "0"
	} else {
		outputString += "1"
	}
	for i := 1; i < n; i++ {
		if zero {
			// If zero is true, start with '0' and alternate
			if i%2 == 0 {
				outputString += "0"
			} else {
				outputString += "1"
			}
		} else {
			// If zero is false, start with '1' and alternate
			if i%2 == 0 {
				outputString += "1"
			} else {
				outputString += "0"
			}
		}
	}

	return outputString
}
