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

// HeapSort implements the heap sort algorithm
// Time Complexity: O(n log n), Space Complexity: O(1)
func HeapSort(array []int) {
	n := len(array)
	
	fmt.Printf("\n%s🏗️  PHASE 1: BUILD MAX HEAP%s\n", Yellow+Bold, Reset)
	fmt.Println("   " + strings.Repeat("─", 45))
	
	// Build max heap (rearrange array)
	// Start from last non-leaf node and heapify each node
	for i := n/2 - 1; i >= 0; i-- {
		fmt.Printf("\n%s📍 Heapifying from index %d:%s\n", Blue+Bold, i, Reset)
		heapify(array, n, i)
		printHeapVisualization(array, n, "Current heap state")
	}
	
	fmt.Printf("\n%s✅ MAX HEAP BUILT!%s Largest element (%d) is at root\n", 
		Green+Bold, Reset, array[0])
	printHeapVisualization(array, n, "Final Max Heap")
	
	fmt.Printf("\n%s🔄 PHASE 2: EXTRACT ELEMENTS%s\n", Magenta+Bold, Reset)
	fmt.Println("   " + strings.Repeat("─", 45))
	
	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		fmt.Printf("\n%s🎯 Step %d:%s Move max element %s%d%s to position %d\n", 
			Cyan+Bold, n-i, Reset, Yellow+Bold, array[0], Reset, i)
		
		// Move current root to end (largest element to its correct position)
		array[0], array[i] = array[i], array[0]
		
		fmt.Printf("   After swap: ")
		printArrayWithSorted(array, i)
		
		// Restore heap property for reduced heap
		fmt.Printf("   %s🔧 Restore heap property for remaining elements%s\n", White, Reset)
		heapify(array, i, 0)
		
		if i > 1 {
			printHeapVisualization(array, i, "Heap after restoration")
		}
	}
}

// heapify ensures the subtree rooted at index i follows max heap property
func heapify(array []int, heapSize, rootIndex int) {
	largest := rootIndex  // Initialize largest as root
	left := 2*rootIndex + 1    // Left child
	right := 2*rootIndex + 2   // Right child
	
	fmt.Printf("      Checking node at index %d (value: %s%d%s)\n", 
		rootIndex, Yellow, array[rootIndex], Reset)
	
	// Check if left child exists and is greater than root
	if left < heapSize {
		fmt.Printf("        Left child [%d]: %s%d%s ", left, Cyan, array[left], Reset)
		if array[left] > array[largest] {
			largest = left
			fmt.Printf("→ %sNew largest!%s\n", Green, Reset)
		} else {
			fmt.Printf("→ Not larger\n")
		}
	}
	
	// Check if right child exists and is greater than current largest
	if right < heapSize {
		fmt.Printf("        Right child [%d]: %s%d%s ", right, Blue, array[right], Reset)
		if array[right] > array[largest] {
			largest = right
			fmt.Printf("→ %sNew largest!%s\n", Green, Reset)
		} else {
			fmt.Printf("→ Not larger\n")
		}
	}
	
	// If largest is not root, swap and continue heapifying
	if largest != rootIndex {
		fmt.Printf("        %s🔄 Swapping:%s %d ↔ %d\n", Red+Bold, Reset, array[rootIndex], array[largest])
		array[rootIndex], array[largest] = array[largest], array[rootIndex]
		
		// Recursively heapify the affected subtree
		heapify(array, heapSize, largest)
	} else {
		fmt.Printf("        %s✅ Heap property satisfied%s\n", Green, Reset)
	}
}

// printHeapVisualization shows the array as a binary tree structure
func printHeapVisualization(array []int, size int, title string) {
	fmt.Printf("\n   %s%s%s %s:\n", Bold, title, Reset, Reset)
	fmt.Printf("   ")
	for i := 0; i < size; i++ {
		if i == 0 {
			fmt.Printf("%s[%d]%s", Yellow+Bold, array[i], Reset) // Root
		} else {
			fmt.Printf(" %s%d%s", Cyan, array[i], Reset)
		}
	}
	fmt.Printf("\n")
	
	// Show tree structure for small arrays
	if size <= 7 {
		fmt.Printf("   Tree view: ")
		if size > 0 {
			fmt.Printf("       %s%d%s\n", Yellow+Bold, array[0], Reset)
		}
		if size > 1 {
			fmt.Printf("              /   \\\n")
			fmt.Printf("            ")
			if size > 1 {
				fmt.Printf("%s%d%s", Green, array[1], Reset)
			}
			fmt.Printf("     ")
			if size > 2 {
				fmt.Printf("%s%d%s", Blue, array[2], Reset)
			}
			fmt.Printf("\n")
		}
		if size > 3 {
			fmt.Printf("           / |   | \\\n")
			fmt.Printf("          ")
			for i := 3; i < 7 && i < size; i++ {
				fmt.Printf("%d ", array[i])
			}
			fmt.Printf("\n")
		}
	}
}

// printArrayWithSorted shows array with sorted portion highlighted
func printArrayWithSorted(array []int, sortedStart int) {
	fmt.Printf("%s[", Cyan)
	for i := 0; i < len(array); i++ {
		if i >= sortedStart {
			fmt.Printf("%s%d%s", Green+Bold, array[i], Cyan) // Sorted portion
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
	originalArray := []int{12, 11, 13, 5, 6, 7}
	
	fmt.Printf("\n%s%s╔══════════════════════════════════════════════╗%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s║          🏔️  HEAP SORT DEMONSTRATION          ║%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════╝%s\n", Bold, Cyan, Reset)
	
	fmt.Printf("\n%s📋 Starting Array:%s %s%v%s\n", Bold+White, Reset, Red+Bold, originalArray, Reset)
	
	fmt.Printf("\n%s%s🚀 Starting Heap Sort Process...%s\n", Bold, Yellow, Reset)
	fmt.Println(strings.Repeat("═", 50))
	
	// Make a copy so we don't modify the original for display
	arrayToSort := make([]int, len(originalArray))
	copy(arrayToSort, originalArray)
	
	// Perform heap sort
	HeapSort(arrayToSort)
	
	fmt.Printf("\n%s%s🎉 FINAL RESULT%s\n", Bold, Green, Reset)
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("   %sOriginal:%s %s%v%s\n", Bold, Reset, Red, originalArray, Reset)
	fmt.Printf("   %sSorted:%s   %s%v%s\n", Bold, Reset, Green+Bold, arrayToSort, Reset)
	
	fmt.Printf("\n%s%s📚 HOW HEAP SORT WORKS%s\n", Bold, Magenta, Reset)
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("   %s1. 🏗️  BUILD HEAP:%s    Convert array into max heap (largest at top)\n", Yellow+Bold, Reset)
	fmt.Printf("   %s2. 🎯 EXTRACT MAX:%s   Move max element to end, reduce heap size\n", Blue+Bold, Reset)
	fmt.Printf("   %s3. 🔧 RESTORE HEAP:%s  Fix heap property for remaining elements\n", Red+Bold, Reset)
	fmt.Printf("   %s4. 🔄 REPEAT:%s       Continue until all elements are sorted\n", Green+Bold, Reset)
	
	fmt.Printf("\n%s⚡ Performance:%s\n", Bold+White, Reset)
	fmt.Printf("   %sTime Complexity:%s   O(n log n) - ALWAYS! No worst case like Quick Sort\n", Cyan, Reset)
	fmt.Printf("   %sSpace Complexity:%s  O(1) - sorts in-place\n", Cyan, Reset)
	fmt.Printf("   %s💡 Advantage:%s      Guaranteed O(n log n), good for real-time systems\n", Green, Reset)
	fmt.Printf("   %s⚠️  Note:%s          Generally slower than Quick Sort in practice\n", Yellow, Reset)
	fmt.Println()
}