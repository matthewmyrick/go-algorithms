# Playlist Manager 🎵🎧

You're building a music streaming app and need to implement a playlist manager! Users can create playlists, add songs, remove songs, and navigate through their music. The playlist should support both sequential and shuffle play modes.

## The Problem

Implement a `PlaylistManager` that manages a collection of songs with the following operations:

1. **AddSong(title, artist, duration)** - Add a song to the end of the playlist
2. **RemoveSong(index)** - Remove song at specific index
3. **GetCurrentSong()** - Get the currently playing song
4. **NextSong()** - Move to next song (with wrap-around)
5. **PreviousSong()** - Move to previous song (with wrap-around)
6. **JumpToSong(index)** - Jump directly to song at index
7. **GetPlaylistLength()** - Get total number of songs
8. **GetAllSongs()** - Get all songs in playlist order
9. **Shuffle()** - Randomly shuffle the playlist
10. **IsEmpty()** - Check if playlist is empty

## Song Structure

```go
type Song struct {
    Title    string
    Artist   string
    Duration int // duration in seconds
}
```

## Example Usage

```go
playlist := NewPlaylistManager()
playlist.AddSong("Bohemian Rhapsody", "Queen", 355)
playlist.AddSong("Hotel California", "Eagles", 391)
playlist.AddSong("Stairway to Heaven", "Led Zeppelin", 482)

current := playlist.GetCurrentSong()  // Bohemian Rhapsody
playlist.NextSong()
current = playlist.GetCurrentSong()   // Hotel California
playlist.Shuffle()                    // Songs now in random order
```

## Your Task

Implement the PlaylistManager struct and all its methods.

**Structure:**
```go
type Song struct {
    Title    string
    Artist   string
    Duration int
}

type PlaylistManager struct {
    // TODO: Define your fields here
}

func NewPlaylistManager() *PlaylistManager {
    // TODO: Initialize and return new playlist manager
}

func (p *PlaylistManager) AddSong(title, artist string, duration int) {
    // TODO: Add song to playlist
}

func (p *PlaylistManager) RemoveSong(index int) bool {
    // TODO: Remove song at index, return true if successful
}

func (p *PlaylistManager) GetCurrentSong() *Song {
    // TODO: Return currently selected song
}

func (p *PlaylistManager) NextSong() *Song {
    // TODO: Move to next song and return it
}

func (p *PlaylistManager) PreviousSong() *Song {
    // TODO: Move to previous song and return it
}

func (p *PlaylistManager) JumpToSong(index int) *Song {
    // TODO: Jump to specific song index
}

func (p *PlaylistManager) GetPlaylistLength() int {
    // TODO: Return number of songs in playlist
}

func (p *PlaylistManager) GetAllSongs() []Song {
    // TODO: Return all songs in current order
}

func (p *PlaylistManager) Shuffle() {
    // TODO: Randomly shuffle the playlist
}

func (p *PlaylistManager) IsEmpty() bool {
    // TODO: Check if playlist is empty
}
```

## Test Cases

The main function will test various playlist operations including:
- Adding and removing songs
- Navigation (next, previous, jump)
- Playlist shuffling
- Edge cases with empty playlists

Rock on! 🤘🎸