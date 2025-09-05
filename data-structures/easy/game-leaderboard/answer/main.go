package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name  string
	Score int
}

type GameLeaderboard struct {
	players map[string]int // player name -> score mapping
	sorted  []Player       // players sorted by score (desc)
	dirty   bool          // flag to track if sorting is needed
}

func NewGameLeaderboard() *GameLeaderboard {
	return &GameLeaderboard{
		players: make(map[string]int),
		sorted:  make([]Player, 0),
		dirty:   false,
	}
}

func (g *GameLeaderboard) AddScore(player string, score int) {
	currentScore, exists := g.players[player]
	
	if !exists {
		// New player
		g.players[player] = score
		g.sorted = append(g.sorted, Player{Name: player, Score: score})
		g.dirty = true
	} else if score > currentScore {
		// Update existing player with higher score
		g.players[player] = score
		// Update in sorted slice
		for i := range g.sorted {
			if g.sorted[i].Name == player {
				g.sorted[i].Score = score
				break
			}
		}
		g.dirty = true
	}
	// If new score is lower, keep existing higher score (no update needed)
}

func (g *GameLeaderboard) ensureSorted() {
	if g.dirty {
		sort.Slice(g.sorted, func(i, j int) bool {
			if g.sorted[i].Score == g.sorted[j].Score {
				return g.sorted[i].Name < g.sorted[j].Name // Stable sort by name
			}
			return g.sorted[i].Score > g.sorted[j].Score
		})
		g.dirty = false
	}
}

func (g *GameLeaderboard) GetTopPlayers(k int) []Player {
	g.ensureSorted()
	
	if k > len(g.sorted) {
		k = len(g.sorted)
	}
	
	result := make([]Player, k)
	copy(result, g.sorted[:k])
	return result
}

func (g *GameLeaderboard) GetPlayerRank(player string) int {
	if _, exists := g.players[player]; !exists {
		return 0 // Player not found
	}
	
	g.ensureSorted()
	
	for i, p := range g.sorted {
		if p.Name == player {
			return i + 1 // 1-indexed rank
		}
	}
	return 0 // Should never reach here if player exists
}

func (g *GameLeaderboard) GetPlayerScore(player string) int {
	score, exists := g.players[player]
	if !exists {
		return 0
	}
	return score
}

func (g *GameLeaderboard) GetTotalPlayers() int {
	return len(g.players)
}

func (g *GameLeaderboard) RemovePlayer(player string) bool {
	if _, exists := g.players[player]; !exists {
		return false
	}
	
	delete(g.players, player)
	
	// Remove from sorted slice
	for i, p := range g.sorted {
		if p.Name == player {
			g.sorted = append(g.sorted[:i], g.sorted[i+1:]...)
			break
		}
	}
	
	return true
}

func main() {
	fmt.Println("🎮 Game Leaderboard Results 🎮\n")

	board := NewGameLeaderboard()
	fmt.Printf("Total players: %d", board.GetTotalPlayers())
	if board.GetTotalPlayers() == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Add players
	board.AddScore("Alice", 1000)
	board.AddScore("Bob", 1500)
	board.AddScore("Charlie", 800)
	board.AddScore("Diana", 1200)

	fmt.Printf("After adding 4 players:\n")
	fmt.Printf("Total players: %d", board.GetTotalPlayers())
	if board.GetTotalPlayers() == 4 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Update existing player (should keep higher score)
	board.AddScore("Alice", 1300)  // Higher score
	board.AddScore("Bob", 1000)    // Lower score, should keep 1500

	// Test that scores were updated correctly
	fmt.Printf("Alice's score after update: %d", board.GetPlayerScore("Alice"))
	if board.GetPlayerScore("Alice") == 1300 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Bob's score (should remain 1500): %d", board.GetPlayerScore("Bob"))
	if board.GetPlayerScore("Bob") == 1500 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Get top players
	fmt.Println("\nTop 3 players:")
	topPlayers := board.GetTopPlayers(3)
	expectedTop3 := []Player{
		{"Bob", 1500},
		{"Alice", 1300},
		{"Diana", 1200},
	}
	
	for i, player := range topPlayers {
		fmt.Printf("%d. %s: %d", i+1, player.Name, player.Score)
		if i < len(expectedTop3) && player.Name == expectedTop3[i].Name && player.Score == expectedTop3[i].Score {
			fmt.Println(" ✅")
		} else {
			fmt.Println(" ❌")
		}
	}

	// Get player ranks and scores
	fmt.Printf("\nPlayer ranks:\n")
	fmt.Printf("Alice rank: %d", board.GetPlayerRank("Alice"))
	if board.GetPlayerRank("Alice") == 2 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Bob rank: %d", board.GetPlayerRank("Bob"))
	if board.GetPlayerRank("Bob") == 1 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Charlie rank: %d", board.GetPlayerRank("Charlie"))
	if board.GetPlayerRank("Charlie") == 4 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("NonExistent rank: %d", board.GetPlayerRank("NonExistent"))
	if board.GetPlayerRank("NonExistent") == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Remove a player
	fmt.Printf("\nRemoving Charlie: %t", board.RemovePlayer("Charlie"))
	if board.RemovePlayer("Charlie") == false { // Already removed, should return false
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Total players after removal: %d", board.GetTotalPlayers())
	if board.GetTotalPlayers() == 3 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Charlie rank after removal: %d", board.GetPlayerRank("Charlie"))
	if board.GetPlayerRank("Charlie") == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	fmt.Println("\nAll tests completed! 🏆")
}