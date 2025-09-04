# Robot Warehouse Sorting - Solution 🤖📦

## Approach

This is an array filtering problem where we need to separate elements based on a condition (even vs odd). We can solve this with a simple iteration through the array.

## Algorithm

1. **Initialize** two empty slices for even and odd numbers
2. **Iterate** through each tracking number in the input
3. **Check** if the number is even using the modulo operator (`%`)
4. **Append** the number to the appropriate slice
5. **Return** both slices

**Key insight:** A number is even if `number % 2 == 0`, otherwise it's odd.

**Time Complexity:** O(n) - single pass through the array
**Space Complexity:** O(n) - storage for the result slices

## Alternative Approaches

1. **Two-pass approach:** First count even/odd numbers to pre-allocate slices, then fill them
2. **In-place partitioning:** Rearrange the original array (but this would lose the requirement to maintain order within groups)
3. **Functional approach:** Use filter functions (if Go had built-in filter)

## Key Learning Points

- **Modulo operator** is essential for even/odd checks
- **Slice operations** in Go are efficient for append operations
- **Multiple return values** in Go make this problem elegant to solve
- **Order preservation** is important in many real-world scenarios