package main

import "fmt"

func CanBrewPotion(recipe, pantry string) bool {
	// TODO: Implement your solution here
	return false
}

func main() {
	// Test your function here
	testCases := [][]string{
		{"love", "evolve"},
		{"magic", "gem"},
		{"aab", "baa"},
		{"aab", "ab"},
		{"", "anything"},
		{"something", ""},
	}

	for i, test := range testCases {
		recipe, pantry := test[0], test[1]
		result := CanBrewPotion(recipe, pantry)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("  Recipe: \"%s\"\n", recipe)
		fmt.Printf("  Pantry: \"%s\"\n", pantry)
		fmt.Printf("  Can brew: %t\n\n", result)
	}
}