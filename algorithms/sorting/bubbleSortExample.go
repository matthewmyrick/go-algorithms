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

// BubbleSort implements the bubble sort algorithm
// Time Complexity: O(n²), Space Complexity: O(1)
func BubbleSort(array []int) {
	n := len(array)
	
	fmt.Printf("\n%s🫧 BUBBLE SORT PROCESS:%s\n", Yellow+Bold, Reset)
	fmt.Println("   " + strings.Repeat("─", 55))
	
	// Traverse through all array elements
	for i := 0; i < n-1; i++ {
		fmt.Printf("\n%s📍 Pass %d:%s Bubbling largest element to position %d\n", 
			Blue+Bold, i+1, Reset, n-1-i)
		
		swapped := false // Flag to optimize - if no swaps, array is sorted
		
		// Show current array state
		fmt.Printf("   Starting array: ")
		printArrayWithBubblePortion(array, n-i, -1, -1)
		
		fmt.Printf("\n%s   🔍 Comparing adjacent elements:%s\n", White, Reset)
		
		// Last i elements are already in place
		for j := 0; j < n-i-1; j++ {
			fmt.Printf("      Position %d vs %d: %s%d%s vs %s%d%s → ", 
				j, j+1, Cyan, array[j], Reset, Magenta, array[j+1], Reset)
			
			// Traverse the array from 0 to n-i-1
			// Swap if the element found is greater than the next element
			if array[j] > array[j+1] {
				fmt.Printf("%s%d > %d%s → %s🔄 SWAP!%s\n", 
					Red, array[j], array[j+1], Reset, Red+Bold, Reset)
				
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
				
				// Show array after swap
				fmt.Printf("         After swap: ")
				printArrayWithBubblePortion(array, n-i, j, j+1)
			} else {
				fmt.Printf("%s%d ≤ %d%s → %sNo swap%s\n", 
					Green, array[j], array[j+1], Reset, Green+Bold, Reset)
			}
		}
		
		// Show array after this pass
		fmt.Printf("\n   %s✅ Pass %d complete:%s Largest element %s%d%s is now in position %d\n", 
			Green+Bold, i+1, Reset, Yellow+Bold, array[n-1-i], Reset, n-1-i)
		fmt.Printf("   After pass %d: ", i+1)
		printArrayWithBubblePortion(array, n-i-1, -1, -1)
		
		// If no swapping happened, array is sorted
		if !swapped {
			fmt.Printf("\n   %s🎉 EARLY TERMINATION:%s No swaps needed - array is sorted!\n", 
				Green+Bold, Reset)
			break
		}
		
		fmt.Println("   " + strings.Repeat("─", 55))
	}
}

// printArrayWithBubblePortion shows array with sorted portion highlighted
func printArrayWithBubblePortion(array []int, unsortedEnd, pos1, pos2 int) {
	fmt.Printf("%s[", Reset)
	for i := 0; i < len(array); i++ {
		if i >= unsortedEnd {
			fmt.Printf("%s%d%s", Green+Bold, array[i], Reset) // Sorted portion (bubbled to end)
		} else if i == pos1 {
			fmt.Printf("%s%d%s", Cyan+Bold, array[i], Reset) // First element in comparison
		} else if i == pos2 {
			fmt.Printf("%s%d%s", Magenta+Bold, array[i], Reset) // Second element in comparison
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
	originalArray := []int{64, 34, 25, 12, 22, 11, 90}
	
	fmt.Printf("\n%s%s╔══════════════════════════════════════════════╗%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s║         🫧 BUBBLE SORT DEMONSTRATION          ║%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════╝%s\n", Bold, Cyan, Reset)
	
	fmt.Printf("\n%s📋 Starting Array:%s %s%v%s\n", Bold+White, Reset, Red+Bold, originalArray, Reset)
	
	fmt.Printf("\n%s%s🚀 Starting Bubble Sort Process...%s\n", Bold, Yellow, Reset)
	fmt.Println(strings.Repeat("═", 55))
	
	// Make a copy so we don't modify the original for display
	arrayToSort := make([]int, len(originalArray))
	copy(arrayToSort, originalArray)
	
	// Perform bubble sort
	BubbleSort(arrayToSort)
	
	fmt.Printf("\n%s%s🎉 FINAL RESULT%s\n", Bold, Green, Reset)
	fmt.Println(strings.Repeat("═", 55))
	fmt.Printf("   %sOriginal:%s %s%v%s\n", Bold, Reset, Red, originalArray, Reset)
	fmt.Printf("   %sSorted:%s   %s%v%s\n", Bold, Reset, Green+Bold, arrayToSort, Reset)
	
	fmt.Printf("\n%s%s📚 HOW BUBBLE SORT WORKS%s\n", Bold, Magenta, Reset)
	fmt.Println(strings.Repeat("═", 55))
	fmt.Printf("   %s1. 🔍 COMPARE:%s      Compare adjacent elements in the array\n", Yellow+Bold, Reset)
	fmt.Printf("   %s2. 🔄 SWAP:%s        Swap them if they're in wrong order\n", Blue+Bold, Reset)
	fmt.Printf("   %s3. 🫧 BUBBLE:%s      Largest element 'bubbles' to the end\n", Green+Bold, Reset)
	fmt.Printf("   %s4. 🔄 REPEAT:%s      Continue with remaining unsorted portion\n", Cyan+Bold, Reset)
	fmt.Printf("   %s5. ✅ OPTIMIZE:%s    Stop early if no swaps occur (array is sorted)\n", Red+Bold, Reset)
	
	fmt.Printf("\n%s⚡ Performance:%s\n", Bold+White, Reset)
	fmt.Printf("   %sBest Case:%s        O(n) - already sorted array with optimization\n", Cyan, Reset)
	fmt.Printf("   %sAverage/Worst:%s    O(n²) - reverse sorted array\n", Cyan, Reset)
	fmt.Printf("   %sSpace Complexity:%s O(1) - sorts in-place\n", Cyan, Reset)
	fmt.Printf("   %s💡 Advantages:%s    Simple to understand, stable sort, adaptive\n", Green, Reset)
	fmt.Printf("   %s⚠️  Disadvantages:%s Very slow for large datasets\n", Yellow, Reset)
	fmt.Printf("   %s🎓 Best for:%s      Teaching, very small datasets, nearly sorted data\n", Blue, Reset)
	fmt.Println()
}