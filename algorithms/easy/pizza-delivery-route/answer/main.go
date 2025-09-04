package main

import (
	"fmt"
	"sort"
)

func OptimizeDeliveryRoute(houses []int) []int {
	if len(houses) == 0 {
		return houses
	}
	
	// Create a copy to avoid modifying the original slice
	result := make([]int, len(houses))
	copy(result, houses)
	
	// Sort the copy in ascending order
	sort.Ints(result)
	
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{42, 15, 8, 23, 4, 16, 35},
		{1},
		{},
		{5, 3, 8, 1, 9, 2},
	}

	expected := [][]int{
		{4, 8, 15, 16, 23, 35, 42},
		{1},
		{},
		{1, 2, 3, 5, 8, 9},
	}

	fmt.println("🍕 Pizza Delivery Route Optimizer Results 🍕\n")

	for i, houses := range testCases {
		result := OptimizeDeliveryRoute(houses)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("  Input:    %v\n", houses)
		fmt.Printf("  Output:   %v\n", result)
		fmt.Printf("  Expected: %v\n", expected[i])
		
		// Check if result matches expected
		if len(result) == len(expected[i]) {
			match := true
			for j := 0; j < len(result); j++ {
				if result[j] != expected[i][j] {
					match = false
					break
				}
			}
			if match {
				fmt.Printf("  Status:   ✅ PASS\n\n")
			} else {
				fmt.Printf("  Status:   ❌ FAIL\n\n")
			}
		} else {
			fmt.Printf("  Status:   ❌ FAIL\n\n")
		}
	}
}