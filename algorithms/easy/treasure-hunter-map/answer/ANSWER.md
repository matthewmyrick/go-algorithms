# Treasure Hunter Map - Solution 🗺️💎

## Approach

This is the classic "House Robber" problem in disguise! We need to find the maximum sum of non-adjacent elements in an array.

## Algorithm - Dynamic Programming

We can solve this using dynamic programming with the following approach:

For each treasure location, we have two choices:
1. **Take it:** Add current treasure to the max treasure from 2 locations back
2. **Skip it:** Keep the max treasure from the previous location

**Recurrence relation:**
```
dp[i] = max(dp[i-1], dp[i-2] + treasures[i])
```

Where:
- `dp[i]` = maximum treasure up to location i
- `dp[i-1]` = max treasure if we skip current location
- `dp[i-2] + treasures[i]` = max treasure if we take current location

**Time Complexity:** O(n)
**Space Complexity:** O(1) - we only need to track the last two values

## Step-by-step Example

For `[2, 7, 9, 3, 1]`:

- Location 0: max = 2 (take 2)
- Location 1: max = max(2, 0+7) = 7 (take 7, skip 2)
- Location 2: max = max(7, 2+9) = 11 (take 2+9, skip 7)
- Location 3: max = max(11, 7+3) = 11 (skip 3, keep previous max)
- Location 4: max = max(11, 11+1) = 12 (take 1 + previous max excluding adjacent)

Wait, let me recalculate:
- dp[0] = 2
- dp[1] = max(2, 7) = 7  
- dp[2] = max(7, 2+9) = 11
- dp[3] = max(11, 7+3) = 11
- dp[4] = max(11, 2+9+1) = 12

The optimal selection is treasures at indices 0, 2, 4 = 2 + 9 + 1 = 12