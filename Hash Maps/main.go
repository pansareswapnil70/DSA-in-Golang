package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hash maps")
	var mymap map[string]string
	fmt.Println("Map using variable type:", mymap)
	var secondMap = make(map[int]int)
	fmt.Println("Map using make function", secondMap)
	ageMap := map[string]string{"Suraj": "Bmt", "Swapnil": "Pune", "Ketan": "Mumbai", "Neel": "Delhi"}
	fmt.Println(ageMap)
	mymap = make(map[string]string)
	mymap["name"] = "Neel"
	mymap["place"] = "Bmt"
	mymap["contact"] = "1234"
	fmt.Println(mymap)
	input := []int{2, 5, 2, 1, 2, 3, 5, 1, 2, 4}
	fmt.Println(input)
	result := recursiveCharacter(input)
	if result != -1 {
		fmt.Println("First Duplicate value found is: ", result)
	} else {
		fmt.Println("No Duplicates Found")
	}
	allDuplicates := findDuplicates(input)
	fmt.Println("All Duplicates found:", allDuplicates)
}

func recursiveCharacter(input []int) int {
	var duplicateCharactersMap = make(map[int]int)
	length := len(input)
	for i := 0; i < length; i++ {
		_, exists := duplicateCharactersMap[input[i]]
		if exists {
			return input[i]
		} else {
			duplicateCharactersMap[input[i]] = 1
		}
	}
	return -1
}

func findDuplicates(input []int) []int {
	var duplicatesMap = make(map[int]int)
	var duplicates []int
	length := len(input)
	for i := 0; i < length; i++ {
		if _, exists := duplicatesMap[input[i]]; exists {
			if duplicatesMap[input[i]] == 1 {
				duplicates = append(duplicates, input[i])
			}
			// Mark the number as fully processed (seen at least twice)
			duplicatesMap[input[i]] = 2
		} else {
			// Mark the number as seen
			duplicatesMap[input[i]] = 1
		}
	}
	return duplicates
}
