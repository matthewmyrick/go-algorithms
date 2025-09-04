package main

import "fmt"

func CanAchieveDockingSequence(arrival, target []int) bool {
	// Edge case: empty sequences
	if len(arrival) != len(target) {
		return false
	}
	if len(arrival) == 0 {
		return true
	}

	var stack []int // holding orbit (stack)
	targetIndex := 0 // current target spacecraft index
	
	// Process each spacecraft in arrival order
	for _, spacecraft := range arrival {
		// Spacecraft enters holding orbit
		stack = append(stack, spacecraft)
		
		// Try to dock spacecraft while possible
		for len(stack) > 0 && targetIndex < len(target) && stack[len(stack)-1] == target[targetIndex] {
			// Spacecraft exits holding orbit and docks
			stack = stack[:len(stack)-1] // pop
			targetIndex++
		}
	}
	
	// Check if all spacecraft have docked successfully
	return len(stack) == 0 && targetIndex == len(target)
}

func main() {
	// Test cases
	testCases := [][][]int{
		{{1, 2, 3, 4, 5}, {3, 2, 4, 1, 5}},
		{{1, 2, 3}, {3, 1, 2}},
		{{1, 2}, {2, 1}},
		{{1}, {1}},
		{{}, {}},
	}

	expected := []bool{true, false, true, true, true}

	fmt.Println("🚀 Space Station Docking Results 🛸\n")

	for i, test := range testCases {
		arrival, target := test[0], test[1]
		result := CanAchieveDockingSequence(arrival, target)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("  Arrival:  %v\n", arrival)
		fmt.Printf("  Target:   %v\n", target)
		fmt.Printf("  Output:   %t\n", result)
		fmt.Printf("  Expected: %t\n", expected[i])
		
		if result == expected[i] {
			fmt.Printf("  Status:   ✅ PASS\n\n")
		} else {
			fmt.Printf("  Status:   ❌ FAIL\n\n")
		}
	}
}