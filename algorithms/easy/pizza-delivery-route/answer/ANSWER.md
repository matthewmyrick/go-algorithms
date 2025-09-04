# Pizza Delivery Route Optimizer - Solution 🍕

## Approach

This is a classic sorting problem. We need to sort the house numbers in ascending order so the delivery driver can visit them efficiently.

## Algorithm

We can use Go's built-in `sort.Ints()` function which implements an efficient sorting algorithm (typically introsort - a hybrid of quicksort, heapsort, and insertion sort).

**Time Complexity:** O(n log n)
**Space Complexity:** O(1) - sorting is done in-place

## Alternative Approaches

1. **Bubble Sort** - O(n²) time, good for learning but inefficient for large inputs
2. **Selection Sort** - O(n²) time, simple but inefficient
3. **Merge Sort** - O(n log n) time, O(n) space, stable sorting
4. **Quick Sort** - O(n log n) average case, O(n²) worst case

## Key Learning Points

- Sorting is fundamental to many algorithms
- Go's standard library provides efficient implementations
- Always consider the trade-offs between time and space complexity
- In real applications, use proven library functions rather than implementing from scratch