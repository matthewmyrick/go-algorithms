package main

import "fmt"

type Song struct {
	Title    string
	Artist   string
	Duration int // duration in seconds
}

type PlaylistManager struct {
	// TODO: Define your fields here
}

func NewPlaylistManager() *PlaylistManager {
	// TODO: Initialize and return new playlist manager
	return nil
}

func (p *PlaylistManager) AddSong(title, artist string, duration int) {
	// TODO: Add song to playlist
}

func (p *PlaylistManager) RemoveSong(index int) bool {
	// TODO: Remove song at index, return true if successful
	return false
}

func (p *PlaylistManager) GetCurrentSong() *Song {
	// TODO: Return currently selected song
	return nil
}

func (p *PlaylistManager) NextSong() *Song {
	// TODO: Move to next song and return it
	return nil
}

func (p *PlaylistManager) PreviousSong() *Song {
	// TODO: Move to previous song and return it
	return nil
}

func (p *PlaylistManager) JumpToSong(index int) *Song {
	// TODO: Jump to specific song index
	return nil
}

func (p *PlaylistManager) GetPlaylistLength() int {
	// TODO: Return number of songs in playlist
	return 0
}

func (p *PlaylistManager) GetAllSongs() []Song {
	// TODO: Return all songs in current order
	return nil
}

func (p *PlaylistManager) Shuffle() {
	// TODO: Randomly shuffle the playlist
}

func (p *PlaylistManager) IsEmpty() bool {
	// TODO: Check if playlist is empty
	return true
}

func formatDuration(seconds int) string {
	minutes := seconds / 60
	secs := seconds % 60
	return fmt.Sprintf("%d:%02d", minutes, secs)
}

func main() {
	// Test your implementation
	fmt.Println("🎵 Testing Playlist Manager 🎵\n")

	playlist := NewPlaylistManager()
	fmt.Printf("Is empty: %t\n", playlist.IsEmpty())
	fmt.Printf("Length: %d\n", playlist.GetPlaylistLength())
	fmt.Printf("Current song: %v\n\n", playlist.GetCurrentSong())

	// Add songs
	playlist.AddSong("Bohemian Rhapsody", "Queen", 355)
	playlist.AddSong("Hotel California", "Eagles", 391)
	playlist.AddSong("Stairway to Heaven", "Led Zeppelin", 482)
	playlist.AddSong("Sweet Child O' Mine", "Guns N' Roses", 356)

	fmt.Printf("After adding 4 songs:\n")
	fmt.Printf("Length: %d\n", playlist.GetPlaylistLength())
	fmt.Printf("Is empty: %t\n", playlist.IsEmpty())

	// Test current song
	current := playlist.GetCurrentSong()
	if current != nil {
		fmt.Printf("Current song: %s by %s (%s)\n", current.Title, current.Artist, formatDuration(current.Duration))
	}

	// Test navigation
	fmt.Println("\nTesting navigation:")
	for i := 0; i < 5; i++ { // Go beyond playlist length to test wrap-around
		next := playlist.NextSong()
		if next != nil {
			fmt.Printf("Next: %s by %s\n", next.Title, next.Artist)
		}
	}

	fmt.Println("\nGoing backwards:")
	for i := 0; i < 3; i++ {
		prev := playlist.PreviousSong()
		if prev != nil {
			fmt.Printf("Previous: %s by %s\n", prev.Title, prev.Artist)
		}
	}

	// Test jump to song
	fmt.Println("\nJumping to song index 2:")
	jumped := playlist.JumpToSong(2)
	if jumped != nil {
		fmt.Printf("Jumped to: %s by %s\n", jumped.Title, jumped.Artist)
	}

	// Test shuffle
	fmt.Println("\nBefore shuffle:")
	allSongs := playlist.GetAllSongs()
	for i, song := range allSongs {
		fmt.Printf("%d. %s by %s\n", i+1, song.Title, song.Artist)
	}

	playlist.Shuffle()
	fmt.Println("\nAfter shuffle:")
	allSongs = playlist.GetAllSongs()
	for i, song := range allSongs {
		fmt.Printf("%d. %s by %s\n", i+1, song.Title, song.Artist)
	}

	// Test remove song
	fmt.Printf("\nRemoving song at index 1: %t\n", playlist.RemoveSong(1))
	fmt.Printf("Length after removal: %d\n", playlist.GetPlaylistLength())

	// Test edge cases
	fmt.Printf("Removing invalid index (-1): %t\n", playlist.RemoveSong(-1))
	fmt.Printf("Removing invalid index (100): %t\n", playlist.RemoveSong(100))
	fmt.Printf("Jumping to invalid index (100): %v\n", playlist.JumpToSong(100))
}