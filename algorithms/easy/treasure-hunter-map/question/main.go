package main

import "fmt"

func MaxTreasure(treasures []int) int {
	// TODO: Implement your solution here
	return 0
}

func main() {
	// Test your function here
	testCases := [][]int{
		{2, 7, 9, 3, 1},
		{5, 3, 4, 11, 2},
		{1, 2, 3},
		{5},
		{},
	}

	for i, treasures := range testCases {
		result := MaxTreasure(treasures)
		fmt.Printf("Test Case %d: %v -> %d\n", i+1, treasures, result)
	}
}