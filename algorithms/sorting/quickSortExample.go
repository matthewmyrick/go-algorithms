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

// QuickSort implements the divide-and-conquer quick sort algorithm
// Time Complexity: O(n log n) average, O(n²) worst case, Space Complexity: O(log n)
func QuickSort(array []int, low, high int) {
	if low < high {
		// Find pivot position and partition the array
		pivotIndex := partition(array, low, high)
		
		fmt.Printf("\n%s🎯 PIVOT PLACED:%s Pivot %s%d%s is now in correct position at index %d\n", 
			Green+Bold, Reset, Yellow+Bold, array[pivotIndex], Reset, pivotIndex)
		fmt.Printf("   Array: %s", Cyan)
		for i := low; i <= high; i++ {
			if i == pivotIndex {
				fmt.Printf("[%s%d%s]", Yellow+Bold, array[i], Cyan)
			} else {
				fmt.Printf("%d", array[i])
			}
			if i < high {
				fmt.Print(" ")
			}
		}
		fmt.Printf("%s\n", Reset)
		fmt.Println("   " + strings.Repeat("─", 50))
		
		// Recursively sort left and right partitions
		if pivotIndex-1 > low {
			fmt.Printf("\n%s📌 Sorting LEFT partition:%s [%d to %d]\n", Blue+Bold, Reset, low, pivotIndex-1)
			QuickSort(array, low, pivotIndex-1)
		}
		
		if pivotIndex+1 < high {
			fmt.Printf("\n%s📌 Sorting RIGHT partition:%s [%d to %d]\n", Red+Bold, Reset, pivotIndex+1, high)
			QuickSort(array, pivotIndex+1, high)
		}
	}
}

// partition rearranges the array so elements smaller than pivot are on left,
// elements greater than pivot are on right, and returns the pivot's final position
func partition(array []int, low, high int) int {
	// Choose rightmost element as pivot
	pivot := array[high]
	fmt.Printf("\n%s🎯 PARTITIONING:%s Pivot = %s%d%s (rightmost element)\n", 
		Magenta+Bold, Reset, Yellow+Bold, pivot, Reset)
	
	// Index of smaller element (indicates right position of pivot)
	i := low - 1
	
	fmt.Printf("%s   🔍 Comparing each element with pivot:%s\n", White, Reset)
	
	for j := low; j < high; j++ {
		fmt.Printf("      Comparing %s%d%s with pivot %s%d%s: ", 
			Cyan, array[j], Reset, Yellow, pivot, Reset)
		
		// If current element is smaller than or equal to pivot
		if array[j] <= pivot {
			i++ // Increment index of smaller element
			array[i], array[j] = array[j], array[i] // Swap elements
			fmt.Printf("%s%d ≤ %d%s → %sSwap%s → ", 
				Green, array[j], pivot, Reset, Green+Bold, Reset)
		} else {
			fmt.Printf("%s%d > %d%s → %sNo swap%s → ", 
				Red, array[j], pivot, Reset, Red+Bold, Reset)
		}
		
		// Show current array state
		fmt.Printf("Array: %s", Cyan)
		for k := low; k <= high; k++ {
			if k == i && i >= low {
				fmt.Printf("[%s%d%s]", Green+Bold, array[k], Cyan)
			} else if k == high {
				fmt.Printf("(%s%d%s)", Yellow+Bold, array[k], Cyan)
			} else {
				fmt.Printf("%d", array[k])
			}
			if k < high {
				fmt.Print(" ")
			}
		}
		fmt.Printf("%s\n", Reset)
	}
	
	// Place pivot in correct position
	array[i+1], array[high] = array[high], array[i+1]
	fmt.Printf("\n%s   🎯 Final step:%s Place pivot in correct position\n", White, Reset)
	
	return i + 1
}

func main() {
	// Hardcoded array for learning purposes
	originalArray := []int{64, 34, 25, 12, 22, 11, 90}
	
	fmt.Printf("\n%s%s╔══════════════════════════════════════════════╗%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s║          🚀 QUICK SORT DEMONSTRATION          ║%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════╝%s\n", Bold, Cyan, Reset)
	
	fmt.Printf("\n%s📋 Starting Array:%s %s%v%s\n", Bold+White, Reset, Red+Bold, originalArray, Reset)
	
	fmt.Printf("\n%s%s🚀 Starting Quick Sort Process...%s\n", Bold, Yellow, Reset)
	fmt.Println(strings.Repeat("═", 50))
	
	// Make a copy so we don't modify the original for display
	arrayToSort := make([]int, len(originalArray))
	copy(arrayToSort, originalArray)
	
	// Perform quick sort
	QuickSort(arrayToSort, 0, len(arrayToSort)-1)
	
	fmt.Printf("\n%s%s🎉 FINAL RESULT%s\n", Bold, Green, Reset)
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("   %sOriginal:%s %s%v%s\n", Bold, Reset, Red, originalArray, Reset)
	fmt.Printf("   %sSorted:%s   %s%v%s\n", Bold, Reset, Green+Bold, arrayToSort, Reset)
	
	fmt.Printf("\n%s%s📚 HOW QUICK SORT WORKS%s\n", Bold, Magenta, Reset)
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("   %s1. 🎯 CHOOSE PIVOT:%s Pick an element (we use rightmost)\n", Yellow+Bold, Reset)
	fmt.Printf("   %s2. 🔄 PARTITION:%s   Rearrange so smaller elements are left, larger are right\n", Blue+Bold, Reset)
	fmt.Printf("   %s3. 📌 RECURSE:%s     Repeat process on left and right partitions\n", Green+Bold, Reset)
	fmt.Printf("   %s4. ✅ COMBINE:%s     No combining needed - array sorts in-place!\n", Cyan+Bold, Reset)
	
	fmt.Printf("\n%s⚡ Performance:%s\n", Bold+White, Reset)
	fmt.Printf("   %sAverage Time:%s    O(n log n) - good pivot choices\n", Cyan, Reset)
	fmt.Printf("   %sWorst Case:%s      O(n²) - already sorted arrays with bad pivot\n", Cyan, Reset)
	fmt.Printf("   %sSpace Complexity:%s O(log n) - recursion stack\n", Cyan, Reset)
	fmt.Printf("   %s💡 Advantage:%s     Sorts in-place, very fast in practice!\n", Green, Reset)
	fmt.Println()
}