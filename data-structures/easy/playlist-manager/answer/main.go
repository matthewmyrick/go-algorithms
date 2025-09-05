package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Song struct {
	Title    string
	Artist   string
	Duration int // duration in seconds
}

type PlaylistManager struct {
	songs        []Song
	currentIndex int
}

func NewPlaylistManager() *PlaylistManager {
	rand.Seed(time.Now().UnixNano())
	return &PlaylistManager{
		songs:        make([]Song, 0),
		currentIndex: 0,
	}
}

func (p *PlaylistManager) AddSong(title, artist string, duration int) {
	song := Song{
		Title:    title,
		Artist:   artist,
		Duration: duration,
	}
	p.songs = append(p.songs, song)
}

func (p *PlaylistManager) RemoveSong(index int) bool {
	if index < 0 || index >= len(p.songs) {
		return false
	}
	
	p.songs = append(p.songs[:index], p.songs[index+1:]...)
	
	// Adjust current index
	if len(p.songs) == 0 {
		p.currentIndex = 0
	} else if index < p.currentIndex {
		p.currentIndex--
	} else if index == p.currentIndex && p.currentIndex >= len(p.songs) {
		p.currentIndex = 0
	}
	
	return true
}

func (p *PlaylistManager) GetCurrentSong() *Song {
	if len(p.songs) == 0 {
		return nil
	}
	return &p.songs[p.currentIndex]
}

func (p *PlaylistManager) NextSong() *Song {
	if len(p.songs) == 0 {
		return nil
	}
	p.currentIndex = (p.currentIndex + 1) % len(p.songs)
	return &p.songs[p.currentIndex]
}

func (p *PlaylistManager) PreviousSong() *Song {
	if len(p.songs) == 0 {
		return nil
	}
	p.currentIndex = (p.currentIndex - 1 + len(p.songs)) % len(p.songs)
	return &p.songs[p.currentIndex]
}

func (p *PlaylistManager) JumpToSong(index int) *Song {
	if index < 0 || index >= len(p.songs) {
		return nil
	}
	p.currentIndex = index
	return &p.songs[p.currentIndex]
}

func (p *PlaylistManager) GetPlaylistLength() int {
	return len(p.songs)
}

func (p *PlaylistManager) GetAllSongs() []Song {
	// Return a copy to prevent external modifications
	result := make([]Song, len(p.songs))
	copy(result, p.songs)
	return result
}

func (p *PlaylistManager) Shuffle() {
	// Fisher-Yates shuffle algorithm
	for i := len(p.songs) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		p.songs[i], p.songs[j] = p.songs[j], p.songs[i]
	}
	p.currentIndex = 0 // Reset to start after shuffle
}

func (p *PlaylistManager) IsEmpty() bool {
	return len(p.songs) == 0
}

func formatDuration(seconds int) string {
	minutes := seconds / 60
	secs := seconds % 60
	return fmt.Sprintf("%d:%02d", minutes, secs)
}

func main() {
	fmt.Println("🎵 Playlist Manager Results 🎵\n")

	playlist := NewPlaylistManager()
	fmt.Printf("Is empty: %t", playlist.IsEmpty())
	if playlist.IsEmpty() {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Length: %d", playlist.GetPlaylistLength())
	if playlist.GetPlaylistLength() == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Current song: %v", playlist.GetCurrentSong())
	if playlist.GetCurrentSong() == nil {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Add songs
	playlist.AddSong("Bohemian Rhapsody", "Queen", 355)
	playlist.AddSong("Hotel California", "Eagles", 391)
	playlist.AddSong("Stairway to Heaven", "Led Zeppelin", 482)
	playlist.AddSong("Sweet Child O' Mine", "Guns N' Roses", 356)

	fmt.Printf("\nAfter adding 4 songs:\n")
	fmt.Printf("Length: %d", playlist.GetPlaylistLength())
	if playlist.GetPlaylistLength() == 4 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Is empty: %t", playlist.IsEmpty())
	if !playlist.IsEmpty() {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test current song
	current := playlist.GetCurrentSong()
	if current != nil {
		fmt.Printf("Current song: %s by %s (%s)", current.Title, current.Artist, formatDuration(current.Duration))
		if current.Title == "Bohemian Rhapsody" {
			fmt.Println(" ✅")
		} else {
			fmt.Println(" ❌")
		}
	}

	// Test navigation
	fmt.Println("\nTesting navigation:")
	expectedSongs := []string{"Hotel California", "Stairway to Heaven", "Sweet Child O' Mine", "Bohemian Rhapsody", "Hotel California"}
	
	for i := 0; i < 5; i++ {
		next := playlist.NextSong()
		if next != nil {
			fmt.Printf("Next: %s", next.Title)
			if next.Title == expectedSongs[i] {
				fmt.Println(" ✅")
			} else {
				fmt.Println(" ❌")
			}
		}
	}

	// Test remove song
	originalLength := playlist.GetPlaylistLength()
	fmt.Printf("\nRemoving song at index 1: %t", playlist.RemoveSong(1))
	if playlist.RemoveSong(1) == false { // Already removed
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Length after removal: %d", playlist.GetPlaylistLength())
	if playlist.GetPlaylistLength() == originalLength-1 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test edge cases
	fmt.Printf("Removing invalid index (-1): %t", playlist.RemoveSong(-1))
	if playlist.RemoveSong(-1) == false {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Removing invalid index (100): %t", playlist.RemoveSong(100))
	if playlist.RemoveSong(100) == false {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Jumping to invalid index (100): %v", playlist.JumpToSong(100))
	if playlist.JumpToSong(100) == nil {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test jump to valid index
	jumped := playlist.JumpToSong(0)
	if jumped != nil {
		fmt.Printf("Jumped to index 0: %s", jumped.Title)
		fmt.Println(" ✅")
	} else {
		fmt.Println("Failed to jump to valid index ❌")
	}

	fmt.Println("\nAll core tests completed! 🎧")
	
	// Show final playlist state
	fmt.Println("\nFinal playlist:")
	allSongs := playlist.GetAllSongs()
	for i, song := range allSongs {
		fmt.Printf("%d. %s by %s (%s)\n", i+1, song.Title, song.Artist, formatDuration(song.Duration))
	}
}