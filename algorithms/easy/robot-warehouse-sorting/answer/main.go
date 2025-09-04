package main

import "fmt"

func SeparatePackages(trackingNumbers []int) ([]int, []int) {
	var evenNumbers []int
	var oddNumbers []int

	for _, num := range trackingNumbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		} else {
			oddNumbers = append(oddNumbers, num)
		}
	}

	return evenNumbers, oddNumbers
}

func main() {
	// Test cases
	testCases := [][]int{
		{12, 7, 18, 3, 9, 14, 5, 20},
		{1, 3, 5, 7},
		{2, 4, 6, 8},
		{42},
		{},
	}

	expectedEven := [][]int{
		{12, 18, 14, 20},
		{},
		{2, 4, 6, 8},
		{42},
		{},
	}

	expectedOdd := [][]int{
		{7, 3, 9, 5},
		{1, 3, 5, 7},
		{},
		{},
		{},
	}

	fmt.Println("🤖 Robot Warehouse Sorting Results 📦\n")

	for i, numbers := range testCases {
		even, odd := SeparatePackages(numbers)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("  Input:         %v\n", numbers)
		fmt.Printf("  Even Output:   %v\n", even)
		fmt.Printf("  Even Expected: %v\n", expectedEven[i])
		fmt.Printf("  Odd Output:    %v\n", odd)
		fmt.Printf("  Odd Expected:  %v\n", expectedOdd[i])
		
		// Check if results match expected
		evenMatch := slicesEqual(even, expectedEven[i])
		oddMatch := slicesEqual(odd, expectedOdd[i])
		
		if evenMatch && oddMatch {
			fmt.Printf("  Status:        ✅ PASS\n\n")
		} else {
			fmt.Printf("  Status:        ❌ FAIL\n\n")
		}
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}