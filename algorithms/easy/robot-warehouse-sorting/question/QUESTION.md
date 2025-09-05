# Robot Warehouse Sorting 🤖📦

You're programming a robot in an automated warehouse! The robot needs to separate packages into two groups based on whether their tracking numbers are even or odd. This helps the warehouse organize packages for different shipping routes.

## The Problem

Given an array of package tracking numbers (integers), separate them into two arrays:
1. **Even numbers** - for the overnight shipping route
2. **Odd numbers** - for the standard shipping route

The robot should maintain the original order of packages within each group.

## Example

**Input:** `[12, 7, 18, 3, 9, 14, 5, 20]`

**Output:** 
```
Even: [12, 18, 14, 20]
Odd:  [7, 3, 9, 5]
```

## Your Task

Write a function that takes a slice of package tracking numbers and returns two slices: one for even numbers and one for odd numbers.

**Function signature:**
```go
func SeparatePackages(trackingNumbers []int) ([]int, []int)
```

The function should return `(evenNumbers, oddNumbers)`.

## Test Cases

1. `[12, 7, 18, 3, 9, 14, 5, 20]` → `([12, 18, 14, 20], [7, 3, 9, 5])`
2. `[1, 3, 5, 7]` → `([], [1, 3, 5, 7])`
3. `[2, 4, 6, 8]` → `([2, 4, 6, 8], [])`
4. `[42]` → `([42], [])`
5. `[]` → `([], [])`

Let's help the robot sort those packages efficiently! 🚛✨