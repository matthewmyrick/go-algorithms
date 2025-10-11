package main

import "fmt"

func OptimizeDeliveryRoute(houses []int) []int {
	// TODO: Implement your solution here
	return nil
}

func main() {
	// Test your function here
	testCases := [][]int{
		{42, 15, 8, 23, 4, 16, 35},
		{1},
		{},
		{5, 3, 8, 1, 9, 2},}

	for i, houses := range testCases {
		result := OptimizeDeliveryRoute(houses)
		fmt.Printf("Test Case %d: %v -> %v\n", i+1, houses, result)
	}
}
