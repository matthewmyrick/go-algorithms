package main

import (
	"fmt"
	"strings"
)

// ANSI color codes for better visualization
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Bold    = "\033[1m"
)

// SelectionSort implements the selection sort algorithm
// Time Complexity: O(n²), Space Complexity: O(1)
func SelectionSort(array []int) {
	n := len(array)
	
	fmt.Printf("\n%s🔍 SELECTION SORT PROCESS:%s\n", Yellow+Bold, Reset)
	fmt.Println("   " + strings.Repeat("─", 50))
	
	// Traverse through all array elements
	for i := 0; i < n-1; i++ {
		fmt.Printf("\n%s📍 Pass %d:%s Finding minimum in unsorted portion [%d to %d]\n", 
			Blue+Bold, i+1, Reset, i, n-1)
		
		// Find the minimum element in remaining unsorted array
		minIndex := i
		fmt.Printf("   🎯 Current minimum: %s%d%s at index %d\n", 
			Yellow+Bold, array[minIndex], Reset, minIndex)
		
		// Show current array state
		fmt.Printf("   Current array: ")
		printArrayWithSortedPortion(array, i, minIndex, -1)
		
		fmt.Printf("\n%s   🔍 Searching for smaller elements:%s\n", White, Reset)
		
		for j := i + 1; j < n; j++ {
			fmt.Printf("      Comparing %s%d%s with current min %s%d%s: ", 
				Cyan, array[j], Reset, Yellow, array[minIndex], Reset)
			
			if array[j] < array[minIndex] {
				minIndex = j
				fmt.Printf("%s%d < %d%s → %sNew minimum found!%s\n", 
					Green, array[j], array[i], Reset, Green+Bold, Reset)
			} else {
				fmt.Printf("%s%d ≥ %d%s → No change\n", 
					Red, array[j], array[minIndex], Reset)
			}
		}
		
		// Swap the found minimum element with the first element
		if minIndex != i {
			fmt.Printf("\n   %s🔄 SWAPPING:%s %s%d%s (index %d) ↔ %s%d%s (index %d)\n", 
				Magenta+Bold, Reset, Yellow+Bold, array[minIndex], Reset, minIndex, 
				Green+Bold, array[i], Reset, i)
			array[i], array[minIndex] = array[minIndex], array[i]
		} else {
			fmt.Printf("\n   %s✅ NO SWAP NEEDED:%s %s%d%s is already in correct position\n", 
				Green+Bold, Reset, Yellow+Bold, array[i], Reset)
		}
		
		// Show array after this pass
		fmt.Printf("   After pass %d: ", i+1)
		printArrayWithSortedPortion(array, i+1, -1, -1)
		fmt.Println("   " + strings.Repeat("─", 50))
	}
}

// printArrayWithSortedPortion shows array with sorted portion highlighted
func printArrayWithSortedPortion(array []int, sortedEnd, currentMin, comparing int) {
	fmt.Printf("%s[", Reset)
	for i := 0; i < len(array); i++ {
		if i < sortedEnd {
			fmt.Printf("%s%d%s", Green+Bold, array[i], Reset) // Sorted portion
		} else if i == currentMin {
			fmt.Printf("%s%d%s", Yellow+Bold, array[i], Reset) // Current minimum
		} else if i == comparing {
			fmt.Printf("%s%d%s", Cyan+Bold, array[i], Reset) // Element being compared
		} else {
			fmt.Printf("%d", array[i]) // Unsorted portion
		}
		if i < len(array)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("]%s\n", Reset)
}

func main() {
	// Hardcoded array for learning purposes
	originalArray := []int{64, 25, 12, 22, 11}
	
	fmt.Printf("\n%s%s╔══════════════════════════════════════════════╗%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s║        🎯 SELECTION SORT DEMONSTRATION        ║%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════╝%s\n", Bold, Cyan, Reset)
	
	fmt.Printf("\n%s📋 Starting Array:%s %s%v%s\n", Bold+White, Reset, Red+Bold, originalArray, Reset)
	
	fmt.Printf("\n%s%s🚀 Starting Selection Sort Process...%s\n", Bold, Yellow, Reset)
	fmt.Println(strings.Repeat("═", 50))
	
	// Make a copy so we don't modify the original for display
	arrayToSort := make([]int, len(originalArray))
	copy(arrayToSort, originalArray)
	
	// Perform selection sort
	SelectionSort(arrayToSort)
	
	fmt.Printf("\n%s%s🎉 FINAL RESULT%s\n", Bold, Green, Reset)
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("   %sOriginal:%s %s%v%s\n", Bold, Reset, Red, originalArray, Reset)
	fmt.Printf("   %sSorted:%s   %s%v%s\n", Bold, Reset, Green+Bold, arrayToSort, Reset)
	
	fmt.Printf("\n%s%s📚 HOW SELECTION SORT WORKS%s\n", Bold, Magenta, Reset)
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("   %s1. 🔍 FIND MIN:%s     Find the smallest element in unsorted portion\n", Yellow+Bold, Reset)
	fmt.Printf("   %s2. 🔄 SWAP:%s        Move it to the beginning of unsorted portion\n", Blue+Bold, Reset)
	fmt.Printf("   %s3. ➡️  ADVANCE:%s     Move boundary between sorted/unsorted portions\n", Green+Bold, Reset)
	fmt.Printf("   %s4. 🔄 REPEAT:%s      Continue until entire array is sorted\n", Cyan+Bold, Reset)
	
	fmt.Printf("\n%s⚡ Performance:%s\n", Bold+White, Reset)
	fmt.Printf("   %sTime Complexity:%s   O(n²) - always makes same number of comparisons\n", Cyan, Reset)
	fmt.Printf("   %sSpace Complexity:%s  O(1) - sorts in-place\n", Cyan, Reset)
	fmt.Printf("   %s💡 Advantages:%s     Simple to understand, minimal swaps (n-1 max)\n", Green, Reset)
	fmt.Printf("   %s⚠️  Disadvantages:%s  Slow for large datasets, not adaptive\n", Yellow, Reset)
	fmt.Printf("   %s🎓 Best for:%s       Educational purposes, small datasets\n", Blue, Reset)
	fmt.Println()
}