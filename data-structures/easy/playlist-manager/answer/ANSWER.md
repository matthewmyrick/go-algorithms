# Playlist Manager - Solution 🎵🎧

## Approach

We use a **slice to store songs** and maintain a **current index** for navigation. This provides efficient access to songs and easy implementation of sequential operations.

## Data Structure Design

```go
type Song struct {
    Title    string
    Artist   string
    Duration int
}

type PlaylistManager struct {
    songs       []Song
    currentIndex int
}
```

## Key Operations

1. **Add/Remove**: Slice append/remove operations
2. **Navigation**: Index manipulation with wrap-around logic
3. **Shuffle**: Fisher-Yates shuffle algorithm
4. **Current tracking**: Handle index updates when songs are removed

## Algorithm Details

### Navigation with Wrap-around
```go
// Next song with wrap-around
func (p *PlaylistManager) NextSong() *Song {
    if len(p.songs) == 0 {
        return nil
    }
    p.currentIndex = (p.currentIndex + 1) % len(p.songs)
    return &p.songs[p.currentIndex]
}

// Previous song with wrap-around  
func (p *PlaylistManager) PreviousSong() *Song {
    if len(p.songs) == 0 {
        return nil
    }
    p.currentIndex = (p.currentIndex - 1 + len(p.songs)) % len(p.songs)
    return &p.songs[p.currentIndex]
}
```

### Shuffle Implementation
We use the **Fisher-Yates shuffle** algorithm for unbiased randomization:
```go
func (p *PlaylistManager) Shuffle() {
    for i := len(p.songs) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        p.songs[i], p.songs[j] = p.songs[j], p.songs[i]
    }
    p.currentIndex = 0 // Reset to start after shuffle
}
```

### Remove Song with Index Adjustment
When removing songs, we need to handle the current index:
```go
func (p *PlaylistManager) RemoveSong(index int) bool {
    if index < 0 || index >= len(p.songs) {
        return false
    }
    
    p.songs = append(p.songs[:index], p.songs[index+1:]...)
    
    // Adjust current index
    if index < p.currentIndex {
        p.currentIndex--
    } else if index == p.currentIndex && p.currentIndex >= len(p.songs) {
        p.currentIndex = 0
    }
    
    return true
}
```

**Time Complexity:**
- AddSong: O(1) amortized
- RemoveSong: O(n) due to slice operations
- Navigation: O(1)
- Shuffle: O(n)
- GetAllSongs: O(n) for copying

**Space Complexity:** O(n) where n is number of songs

## Alternative Approaches

1. **Doubly Linked List**: Better for frequent insertions/deletions
2. **Circular Linked List**: Natural wrap-around navigation
3. **Deque**: Efficient front/back operations

The slice approach is simple and efficient for typical playlist sizes!