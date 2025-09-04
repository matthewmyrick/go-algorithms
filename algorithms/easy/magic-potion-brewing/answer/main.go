package main

import "fmt"

func CanBrewPotion(recipe, pantry string) bool {
	// Empty recipe can always be brewed
	if len(recipe) == 0 {
		return true
	}
	
	// Non-empty recipe but empty pantry cannot be brewed
	if len(pantry) == 0 {
		return false
	}

	// Count ingredients needed from recipe
	needed := make(map[rune]int)
	for _, ingredient := range recipe {
		needed[ingredient]++
	}

	// Check if pantry has enough ingredients
	for _, ingredient := range pantry {
		if count, exists := needed[ingredient]; exists && count > 0 {
			needed[ingredient]--
		}
	}

	// Check if all ingredients are satisfied
	for _, count := range needed {
		if count > 0 {
			return false
		}
	}

	return true
}

func main() {
	// Test cases
	testCases := [][]string{
		{"love", "evolve"},
		{"magic", "gem"},
		{"aab", "baa"},
		{"aab", "ab"},
		{"", "anything"},
		{"something", ""},
	}

	expected := []bool{true, false, true, false, true, false}

	fmt.Println("🧙‍♂️ Magic Potion Brewing Results ⚗️\n")

	for i, test := range testCases {
		recipe, pantry := test[0], test[1]
		result := CanBrewPotion(recipe, pantry)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("  Recipe:   \"%s\"\n", recipe)
		fmt.Printf("  Pantry:   \"%s\"\n", pantry)
		fmt.Printf("  Output:   %t\n", result)
		fmt.Printf("  Expected: %t\n", expected[i])
		
		if result == expected[i] {
			fmt.Printf("  Status:   ✅ PASS\n\n")
		} else {
			fmt.Printf("  Status:   ❌ FAIL\n\n")
		}
	}
}