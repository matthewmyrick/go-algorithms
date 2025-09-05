# Space Station Docking Sequence 🚀🛸

You're a space traffic controller managing spacecraft docking at the International Space Station! Each spacecraft has a unique ID number, and they're approaching the station in a specific order. However, due to space weather, some spacecraft might arrive out of sequence.

Your job is to determine if the current arrival sequence can be rearranged to match the expected docking sequence using a temporary holding orbit.

## The Problem

You have a **holding orbit** that works like a stack (Last In, First Out). Spacecraft can:
1. **Enter** the holding orbit from the arrival queue
2. **Exit** the holding orbit to dock at the station

Given the arrival order and the target docking order, determine if it's possible to achieve the target sequence using the holding orbit.

## Example

**Arrival Order:** `[1, 2, 3, 4, 5]`
**Target Docking:** `[3, 2, 4, 1, 5]`

**Solution:**
1. Ship 1 enters holding orbit: stack = [1]
2. Ship 2 enters holding orbit: stack = [1, 2]  
3. Ship 3 enters holding orbit: stack = [1, 2, 3]
4. Ship 3 exits and docks: stack = [1, 2], docked = [3]
5. Ship 2 exits and docks: stack = [1], docked = [3, 2]
6. Ship 4 enters holding orbit: stack = [1, 4]
7. Ship 4 exits and docks: stack = [1], docked = [3, 2, 4]
8. Ship 1 exits and docks: stack = [], docked = [3, 2, 4, 1]
9. Ship 5 enters and exits immediately: docked = [3, 2, 4, 1, 5]

**Result:** ✅ **Possible!**

## Your Task

Write a function that determines if the target docking sequence is achievable.

**Function signature:**
```go
func CanAchieveDockingSequence(arrival, target []int) bool
```

## Test Cases

1. `([1, 2, 3, 4, 5], [3, 2, 4, 1, 5])` → `true`
2. `([1, 2, 3], [3, 1, 2])` → `false`
3. `([1, 2], [2, 1])` → `true`  
4. `([1], [1])` → `true`
5. `([], [])` → `true`

Navigate those spacecraft safely to dock! 🌌⭐