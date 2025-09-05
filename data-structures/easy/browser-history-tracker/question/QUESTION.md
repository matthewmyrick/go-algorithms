# Browser History Tracker 🌐🔙

You're building a web browser that needs to track user navigation history! Users can visit new pages, go back to previous pages, and go forward again. Your task is to implement a browser history system using appropriate data structures.

## The Problem

Implement a `BrowserHistory` struct with the following operations:
1. **Visit(url)** - Visit a new URL, clearing any forward history
2. **Back(steps)** - Go back in history by the given number of steps
3. **Forward(steps)** - Go forward in history by the given number of steps
4. **CurrentURL()** - Return the current page URL

## Example

```go
browser := NewBrowserHistory("google.com")
browser.Visit("youtube.com")     // History: [google.com, youtube.com*]
browser.Visit("facebook.com")    // History: [google.com, youtube.com, facebook.com*]
browser.Back(1)                  // History: [google.com, youtube.com*, facebook.com]
browser.Visit("twitter.com")     // History: [google.com, youtube.com, twitter.com*]
browser.Forward(1)               // Can't go forward (returns current)
browser.Back(2)                  // History: [google.com*, youtube.com, twitter.com]
browser.CurrentURL()             // Returns: "google.com"
```

## Your Task

Implement the BrowserHistory struct and its methods.

**Structure:**
```go
type BrowserHistory struct {
    // TODO: Define your fields here
}

func NewBrowserHistory(homepage string) *BrowserHistory {
    // TODO: Initialize and return new BrowserHistory
}

func (b *BrowserHistory) Visit(url string) {
    // TODO: Visit new URL
}

func (b *BrowserHistory) Back(steps int) string {
    // TODO: Go back and return current URL
}

func (b *BrowserHistory) Forward(steps int) string {
    // TODO: Go forward and return current URL
}

func (b *BrowserHistory) CurrentURL() string {
    // TODO: Return current URL
}
```

## Test Cases

The main function will test various navigation scenarios to ensure your implementation handles:
- Basic navigation (visit, back, forward)
- Edge cases (going back/forward too many steps)
- Clearing forward history when visiting new pages

Happy browsing! 🖥️✨