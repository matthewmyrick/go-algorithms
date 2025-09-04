# Pizza Delivery Route Optimizer 🍕

You work at Mario's Pizza Palace, the busiest pizzeria in town! Your job is to help the delivery drivers find the shortest route to deliver pizzas to customers.

## The Problem

Given a list of house numbers on the same street where pizzas need to be delivered, you need to sort them in ascending order so the driver can deliver them efficiently without backtracking.

## Example

**Input:** `[42, 15, 8, 23, 4, 16, 35]`
**Output:** `[4, 8, 15, 16, 23, 35, 42]`

The driver starts from the pizza shop at the beginning of the street (house number 0) and should visit houses in order: 4 → 8 → 15 → 16 → 23 → 35 → 42.

## Your Task

Write a function that takes a slice of house numbers and returns them sorted in ascending order.

**Function signature:**
```go
func OptimizeDeliveryRoute(houses []int) []int
```

## Test Cases

1. `[42, 15, 8, 23, 4, 16, 35]` → `[4, 8, 15, 16, 23, 35, 42]`
2. `[1]` → `[1]`
3. `[]` → `[]`
4. `[5, 3, 8, 1, 9, 2]` → `[1, 2, 3, 5, 8, 9]`

Good luck, and remember: happy customers mean happy tips! 💰