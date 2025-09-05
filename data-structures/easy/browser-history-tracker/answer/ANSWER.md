# Browser History Tracker - Solution 🌐🔙

## Approach

We use a **slice-based approach** to store the history and maintain a current position index. This allows efficient back/forward navigation.

## Data Structure Design

```go
type BrowserHistory struct {
    history []string  // Store all visited URLs
    current int       // Index of current position
}
```

## Key Operations

1. **Visit**: Add URL at current+1, truncate remaining history
2. **Back**: Move current index backward (with bounds checking)
3. **Forward**: Move current index forward (with bounds checking)
4. **CurrentURL**: Return URL at current index

## Algorithm Details

### Visit(url)
- Truncate history after current position (clear forward history)
- Append new URL
- Update current to new position

### Back(steps)
- Calculate new position: `max(0, current - steps)`
- Update current index
- Return URL at new position

### Forward(steps)
- Calculate new position: `min(len(history)-1, current + steps)`
- Update current index
- Return URL at new position

**Time Complexity:**
- Visit: O(1) amortized (slice append)
- Back/Forward: O(1)
- CurrentURL: O(1)

**Space Complexity:** O(n) where n is number of pages visited

## Alternative Approaches

1. **Two Stacks**: One for back history, one for forward history
2. **Doubly Linked List**: Each node represents a page
3. **Circular Buffer**: For limited history size

The slice approach is simple and efficient for this use case!