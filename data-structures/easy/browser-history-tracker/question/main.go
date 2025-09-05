package main

import "fmt"

type BrowserHistory struct {
	// TODO: Define your fields here
}

func NewBrowserHistory(homepage string) *BrowserHistory {
	// TODO: Initialize and return new BrowserHistory
	return nil
}

func (b *BrowserHistory) Visit(url string) {
	// TODO: Visit new URL
}

func (b *BrowserHistory) Back(steps int) string {
	// TODO: Go back and return current URL
	return ""
}

func (b *BrowserHistory) Forward(steps int) string {
	// TODO: Go forward and return current URL
	return ""
}

func (b *BrowserHistory) CurrentURL() string {
	// TODO: Return current URL
	return ""
}

func main() {
	// Test your implementation
	fmt.Println("🌐 Testing Browser History 🌐\n")

	browser := NewBrowserHistory("google.com")
	fmt.Printf("Start at: %s\n", browser.CurrentURL())

	browser.Visit("youtube.com")
	fmt.Printf("Visit youtube.com: %s\n", browser.CurrentURL())

	browser.Visit("facebook.com")
	fmt.Printf("Visit facebook.com: %s\n", browser.CurrentURL())

	fmt.Printf("Back 1 step: %s\n", browser.Back(1))
	fmt.Printf("Back 1 step: %s\n", browser.Back(1))

	browser.Visit("twitter.com")
	fmt.Printf("Visit twitter.com: %s\n", browser.CurrentURL())

	fmt.Printf("Back 2 steps: %s\n", browser.Back(2))
	fmt.Printf("Forward 1 step: %s\n", browser.Forward(1))

	// Test edge cases
	fmt.Printf("\nBack 10 steps (too many): %s\n", browser.Back(10))
	fmt.Printf("Forward 10 steps (too many): %s\n", browser.Forward(10))
}