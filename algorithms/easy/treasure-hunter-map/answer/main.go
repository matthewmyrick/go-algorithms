package main

import "fmt"

func MaxTreasure(treasures []int) int {
	if len(treasures) == 0 {
		return 0
	}
	if len(treasures) == 1 {
		return treasures[0]
	}
	if len(treasures) == 2 {
		return max(treasures[0], treasures[1])
	}

	// dp[i] represents max treasure up to index i
	prev2 := treasures[0]           // dp[i-2]
	prev1 := max(treasures[0], treasures[1]) // dp[i-1]

	for i := 2; i < len(treasures); i++ {
		current := max(prev1, prev2+treasures[i])
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := [][]int{
		{2, 7, 9, 3, 1},
		{5, 3, 4, 11, 2},
		{1, 2, 3},
		{5},
		{},
	}

	expected := []int{12, 16, 4, 5, 0}

	fmt.Println("🗺️ Treasure Hunter Map Results 💎\n")

	for i, treasures := range testCases {
		result := MaxTreasure(treasures)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("  Input:    %v\n", treasures)
		fmt.Printf("  Output:   %d\n", result)
		fmt.Printf("  Expected: %d\n", expected[i])
		
		if result == expected[i] {
			fmt.Printf("  Status:   ✅ PASS\n\n")
		} else {
			fmt.Printf("  Status:   ❌ FAIL\n\n")
		}
	}
}