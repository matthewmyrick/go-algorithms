# Game Leaderboard - Solution 🎮🏆

## Approach

We use a **map for fast player lookups** combined with a **sorted slice for rankings**. This hybrid approach balances lookup speed with ranking efficiency.

## Data Structure Design

```go
type Player struct {
    Name  string
    Score int
}

type GameLeaderboard struct {
    players map[string]int  // player name -> score mapping
    sorted  []Player        // players sorted by score (desc)
    dirty   bool           // flag to track if sorting is needed
}
```

## Key Operations

1. **AddScore**: Update map, mark as dirty
2. **GetTopPlayers**: Sort if dirty, return top K
3. **GetPlayerRank**: Sort if dirty, find player position
4. **Lazy Sorting**: Only sort when rankings are requested

## Algorithm Details

### AddScore(player, score)
- Check if player exists in map
- Update score only if new score is higher
- Mark leaderboard as dirty (needs re-sorting)
- Time: O(1)

### GetTopPlayers(k) & GetPlayerRank(player)
- If dirty, sort the players slice by score (descending)
- Return top k players or find player's position
- Time: O(n log n) when sorting needed, O(k) or O(n) otherwise

### Sorting Strategy
We use lazy evaluation - only sort when ranking information is requested:
```go
if g.dirty {
    sort.Slice(g.sorted, func(i, j int) bool {
        return g.sorted[i].Score > g.sorted[j].Score
    })
    g.dirty = false
}
```

**Time Complexity:**
- AddScore: O(1)
- GetTopPlayers: O(n log n) worst case, O(k) if already sorted
- GetPlayerRank: O(n log n) worst case, O(n) if already sorted
- GetPlayerScore: O(1)
- RemovePlayer: O(1) for map, O(n) for slice

**Space Complexity:** O(n) where n is number of players

## Alternative Approaches

1. **Always sorted**: Keep slice always sorted (expensive inserts)
2. **Heap-based**: Use max-heap for top-K queries
3. **Balanced BST**: For frequent rank queries
4. **Segment tree**: For advanced ranking features

The lazy sorting approach works well for games where scores are updated frequently but rankings are checked less often!