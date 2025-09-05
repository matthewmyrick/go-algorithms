package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

type GameLeaderboard struct {
	// TODO: Define your fields here
}

func NewGameLeaderboard() *GameLeaderboard {
	// TODO: Initialize and return new leaderboard
	return nil
}

func (g *GameLeaderboard) AddScore(player string, score int) {
	// TODO: Add or update player score (keep highest)
}

func (g *GameLeaderboard) GetTopPlayers(k int) []Player {
	// TODO: Return top K players by score
	return nil
}

func (g *GameLeaderboard) GetPlayerRank(player string) int {
	// TODO: Return player's rank (1-indexed), 0 if not found
	return 0
}

func (g *GameLeaderboard) GetPlayerScore(player string) int {
	// TODO: Return player's score, 0 if not found
	return 0
}

func (g *GameLeaderboard) GetTotalPlayers() int {
	// TODO: Return total number of players
	return 0
}

func (g *GameLeaderboard) RemovePlayer(player string) bool {
	// TODO: Remove player, return true if removed
	return false
}

func main() {
	// Test your implementation
	fmt.Println("🎮 Testing Game Leaderboard 🎮\n")

	board := NewGameLeaderboard()
	fmt.Printf("Total players: %d\n", board.GetTotalPlayers())

	// Add players
	board.AddScore("Alice", 1000)
	board.AddScore("Bob", 1500)
	board.AddScore("Charlie", 800)
	board.AddScore("Diana", 1200)

	fmt.Printf("After adding 4 players:\n")
	fmt.Printf("Total players: %d\n", board.GetTotalPlayers())

	// Update existing player (should keep higher score)
	board.AddScore("Alice", 1300)  // Higher score
	board.AddScore("Bob", 1000)    // Lower score, should keep 1500

	// Get top players
	fmt.Println("\nTop 3 players:")
	topPlayers := board.GetTopPlayers(3)
	for i, player := range topPlayers {
		fmt.Printf("%d. %s: %d\n", i+1, player.Name, player.Score)
	}

	// Get player ranks and scores
	fmt.Printf("\nPlayer stats:\n")
	fmt.Printf("Alice rank: %d, score: %d\n", board.GetPlayerRank("Alice"), board.GetPlayerScore("Alice"))
	fmt.Printf("Bob rank: %d, score: %d\n", board.GetPlayerRank("Bob"), board.GetPlayerScore("Bob"))
	fmt.Printf("Charlie rank: %d, score: %d\n", board.GetPlayerRank("Charlie"), board.GetPlayerScore("Charlie"))
	fmt.Printf("NonExistent rank: %d, score: %d\n", board.GetPlayerRank("NonExistent"), board.GetPlayerScore("NonExistent"))

	// Remove a player
	fmt.Printf("\nRemoving Charlie: %t\n", board.RemovePlayer("Charlie"))
	fmt.Printf("Total players after removal: %d\n", board.GetTotalPlayers())
	fmt.Printf("Charlie rank after removal: %d\n", board.GetPlayerRank("Charlie"))

	// Test edge cases
	fmt.Println("\nTop 10 players (more than available):")
	allPlayers := board.GetTopPlayers(10)
	for i, player := range allPlayers {
		fmt.Printf("%d. %s: %d\n", i+1, player.Name, player.Score)
	}
}