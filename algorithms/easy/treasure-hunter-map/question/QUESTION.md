# Treasure Hunter Map 🗺️💎

You're a treasure hunter who has discovered an ancient map! The map shows different locations with treasure values, but you can only visit each location once. Your goal is to find the maximum treasure value you can collect.

## The Problem

Given an array of treasure values at different locations, find the maximum sum you can achieve by selecting any number of treasures, but you **cannot select treasures from adjacent locations** (the ancient curse prevents this!).

## Example

**Input:** `[2, 7, 9, 3, 1]`

**Explanation:**
- Location 0: 2 gold coins
- Location 1: 7 gold coins  
- Location 2: 9 gold coins
- Location 3: 3 gold coins
- Location 4: 1 gold coin

**Optimal selection:** Take treasures at locations 1 and 3 (7 + 3 = 10 gold coins)
**Alternative:** Take treasures at locations 0, 2, and 4 (2 + 9 + 1 = 12 gold coins) ✅

**Output:** `12`

## Your Task

Write a function that takes a slice of treasure values and returns the maximum treasure you can collect without taking from adjacent locations.

**Function signature:**
```go
func MaxTreasure(treasures []int) int
```

## Test Cases

1. `[2, 7, 9, 3, 1]` → `12` (take treasures at indices 0, 2, 4)
2. `[5, 3, 4, 11, 2]` → `16` (take treasures at indices 0, 3)
3. `[1, 2, 3]` → `4` (take treasures at indices 0, 2)
4. `[5]` → `5`
5. `[]` → `0`

May your treasure hunting adventures be profitable! ⚱️✨