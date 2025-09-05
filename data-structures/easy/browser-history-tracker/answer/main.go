package main

import "fmt"

type BrowserHistory struct {
	history []string
	current int
}

func NewBrowserHistory(homepage string) *BrowserHistory {
	return &BrowserHistory{
		history: []string{homepage},
		current: 0,
	}
}

func (b *BrowserHistory) Visit(url string) {
	// Clear forward history by truncating
	b.history = b.history[:b.current+1]
	// Add new URL
	b.history = append(b.history, url)
	// Move to new position
	b.current++
}

func (b *BrowserHistory) Back(steps int) string {
	// Move back by steps, but not before index 0
	b.current = max(0, b.current-steps)
	return b.history[b.current]
}

func (b *BrowserHistory) Forward(steps int) string {
	// Move forward by steps, but not beyond last index
	b.current = min(len(b.history)-1, b.current+steps)
	return b.history[b.current]
}

func (b *BrowserHistory) CurrentURL() string {
	return b.history[b.current]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("🌐 Browser History Tracker Results 🌐\n")

	browser := NewBrowserHistory("google.com")
	fmt.Printf("Start at: %s ✓\n", browser.CurrentURL())

	browser.Visit("youtube.com")
	fmt.Printf("Visit youtube.com: %s", browser.CurrentURL())
	if browser.CurrentURL() == "youtube.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	browser.Visit("facebook.com")
	fmt.Printf("Visit facebook.com: %s", browser.CurrentURL())
	if browser.CurrentURL() == "facebook.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	result := browser.Back(1)
	fmt.Printf("Back 1 step: %s", result)
	if result == "youtube.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	result = browser.Back(1)
	fmt.Printf("Back 1 step: %s", result)
	if result == "google.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	browser.Visit("twitter.com")
	fmt.Printf("Visit twitter.com: %s", browser.CurrentURL())
	if browser.CurrentURL() == "twitter.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	result = browser.Back(2)
	fmt.Printf("Back 2 steps: %s", result)
	if result == "google.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	result = browser.Forward(1)
	fmt.Printf("Forward 1 step: %s", result)
	if result == "twitter.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test edge cases
	fmt.Println("\nEdge Cases:")
	result = browser.Back(10)
	fmt.Printf("Back 10 steps (too many): %s", result)
	if result == "google.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	result = browser.Forward(10)
	fmt.Printf("Forward 10 steps (too many): %s", result)
	if result == "twitter.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test that forward history is cleared on new visit
	browser.Back(1)
	browser.Visit("reddit.com")
	result = browser.Forward(1)
	fmt.Printf("\nAfter new visit, forward returns current: %s", result)
	if result == "reddit.com" {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
}