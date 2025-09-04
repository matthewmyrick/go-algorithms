package main

import "fmt"

func SeparatePackages(trackingNumbers []int) ([]int, []int) {
	// TODO: Implement your solution here
	return nil, nil
}

func main() {
	// Test your function here
	testCases := [][]int{
		{12, 7, 18, 3, 9, 14, 5, 20},
		{1, 3, 5, 7},
		{2, 4, 6, 8},
		{42},
		{},
	}

	for i, numbers := range testCases {
		even, odd := SeparatePackages(numbers)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("  Input: %v\n", numbers)
		fmt.Printf("  Even:  %v\n", even)
		fmt.Printf("  Odd:   %v\n\n", odd)
	}
}