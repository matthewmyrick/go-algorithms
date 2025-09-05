# Game Leaderboard 🎮🏆

You're building a leaderboard system for an online game! Players achieve high scores and you need to maintain a ranked list of the top performers. The leaderboard should efficiently handle score updates and provide quick access to rankings.

## The Problem

Implement a `GameLeaderboard` that manages player scores and rankings. The system should support:

1. **AddScore(player, score)** - Add or update a player's score (keep highest)
2. **GetTopPlayers(k)** - Get the top K players with highest scores
3. **GetPlayerRank(player)** - Get a specific player's current rank (1-indexed)
4. **GetPlayerScore(player)** - Get a specific player's current score
5. **GetTotalPlayers()** - Get the total number of players on the leaderboard
6. **RemovePlayer(player)** - Remove a player from the leaderboard

## Player Structure

```go
type Player struct {
    Name  string
    Score int
}
```

## Example Usage

```go
board := NewGameLeaderboard()
board.AddScore("Alice", 1000)
board.AddScore("Bob", 1500)
board.AddScore("Charlie", 800)
board.AddScore("Alice", 1200)  // Alice's score updated to 1200

fmt.Println(board.GetTopPlayers(2))  // [{Bob 1500} {Alice 1200}]
fmt.Println(board.GetPlayerRank("Alice"))  // 2
fmt.Println(board.GetPlayerScore("Charlie"))  // 800
```

## Your Task

Implement the GameLeaderboard struct and all its methods.

**Structure:**
```go
type Player struct {
    Name  string
    Score int
}

type GameLeaderboard struct {
    // TODO: Define your fields here
}

func NewGameLeaderboard() *GameLeaderboard {
    // TODO: Initialize and return new leaderboard
}

func (g *GameLeaderboard) AddScore(player string, score int) {
    // TODO: Add or update player score (keep highest)
}

func (g *GameLeaderboard) GetTopPlayers(k int) []Player {
    // TODO: Return top K players by score
}

func (g *GameLeaderboard) GetPlayerRank(player string) int {
    // TODO: Return player's rank (1-indexed), 0 if not found
}

func (g *GameLeaderboard) GetPlayerScore(player string) int {
    // TODO: Return player's score, 0 if not found
}

func (g *GameLeaderboard) GetTotalPlayers() int {
    // TODO: Return total number of players
}

func (g *GameLeaderboard) RemovePlayer(player string) bool {
    // TODO: Remove player, return true if removed
}
```

## Test Cases

The main function will test various leaderboard operations including:
- Adding and updating player scores
- Retrieving top players in ranked order
- Getting individual player ranks and scores
- Removing players from the leaderboard

Game on! 🎯🌟