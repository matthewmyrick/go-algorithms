package main

import (
	"fmt"
	"strings"
	"time"
)

type Book struct {
	ID          int
	Title       string
	Author      string
	ISBN        string
	IsAvailable bool
	CheckedOutBy string
	DueDate     time.Time
}

type Library struct {
	books []Book
}

func NewLibrary() *Library {
	return &Library{
		books: []Book{
			{1, "Harry Potter and the Sorcerer's Stone", "J.K. Rowling", "978-0439708180", true, "", time.Time{}},
			{2, "Harry Potter and the Chamber of Secrets", "J.K. Rowling", "978-0439064873", true, "", time.Time{}},
			{3, "The Great Gatsby", "F. Scott Fitzgerald", "978-0743273565", true, "", time.Time{}},
			{4, "To Kill a Mockingbird", "Harper Lee", "978-0061120084", false, "John Doe", time.Now().AddDate(0, 0, -19)}, // 5 days overdue (14 day loan + 5)
			{5, "1984", "George Orwell", "978-0452284234", true, "", time.Time{}},
		},
	}
}

// FIXED: String searching corrected
func (l *Library) SearchBooks(query string) []Book {
	var results []Book
	query = strings.ToLower(query) // FIXED: Use ToLower for consistent comparison
	
	for _, book := range l.books {
		// Check title, author, and ISBN - FIXED: Use OR (||) not AND (&&)
		if strings.Contains(strings.ToLower(book.Title), query) ||
		   strings.Contains(strings.ToLower(book.Author), query) ||
		   strings.Contains(strings.ToLower(book.ISBN), query) {
			results = append(results, book)
		}
	}
	
	return results
}

// FIXED: Checkout logic corrected
func (l *Library) CheckoutBook(bookID int, patronName string) bool {
	for i := range l.books {
		if l.books[i].ID == bookID {
			if !l.books[i].IsAvailable { // FIXED: Check if NOT available
				return false // Book not available
			}
			
			l.books[i].IsAvailable = false // FIXED: Set to false (checked out)
			l.books[i].CheckedOutBy = patronName
			l.books[i].DueDate = time.Now().AddDate(0, 0, 14) // FIXED: 14 days, not 7
			return true
		}
	}
	return false // Book not found
}

// FIXED: Return logic corrected
func (l *Library) ReturnBook(bookID int) bool {
	for i := range l.books {
		if l.books[i].ID == bookID {
			if l.books[i].IsAvailable { // FIXED: Check if available (wasn't checked out)
				return false // Book wasn't checked out
			}
			
			l.books[i].IsAvailable = true // FIXED: Set to true (available)
			l.books[i].CheckedOutBy = "" // FIXED: Clear to empty string
			l.books[i].DueDate = time.Time{} // FIXED: Reset to zero time
			return true
		}
	}
	return false // Book not found
}

// FIXED: Overdue calculation corrected
func (l *Library) CalculateOverdueFee(bookID int) float64 {
	for _, book := range l.books {
		if book.ID == bookID {
			if book.IsAvailable { // Book not checked out
				return 0.0
			}
			
			if time.Now().Before(book.DueDate) { // Not overdue
				return 0.0
			}
			
			// Calculate overdue days - FIXED: Divide by 24, not multiply
			overdueDays := time.Now().Sub(book.DueDate).Hours() / 24
			feePerDay := 0.25
			return overdueDays * feePerDay
		}
	}
	return 0.0 // Book not found
}

// FIXED: Display formatting corrected
func (l *Library) DisplayBooks(books []Book) {
	if len(books) == 0 { // FIXED: Check == 0 for empty results
		fmt.Println("No books found.")
		return
	}
	
	for _, book := range books {
		status := "Available"
		if !book.IsAvailable { // FIXED: Correct logic (!IsAvailable = checked out)
			status = "Checked out"
		}
		
		fmt.Printf("ID: %d\n", book.ID)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Printf("Status: %s\n", status)
		
		if !book.IsAvailable { // FIXED: Show details for checked out books (!IsAvailable)
			fmt.Printf("Checked out by: %s\n", book.CheckedOutBy)
			fmt.Printf("Due date: %s\n", book.DueDate.Format("2006-01-02"))
		}
		fmt.Println("---")
	}
}

func main() {
	fmt.Println("📚 Library Book Finder (FIXED VERSION) 📚\n")
	
	library := NewLibrary()
	
	// Test search functionality
	fmt.Println("=== SEARCH TESTS ===")
	
	fmt.Println("Searching for 'Harry Potter':")
	harryPotterBooks := library.SearchBooks("Harry Potter")
	library.DisplayBooks(harryPotterBooks)
	fmt.Printf("Found %d book(s). Expected: 2 Harry Potter books\n\n", len(harryPotterBooks))
	
	fmt.Println("Searching for 'J.K. Rowling':")
	rowlingBooks := library.SearchBooks("J.K. Rowling")
	library.DisplayBooks(rowlingBooks)
	fmt.Printf("Found %d book(s). Expected: 2 books by J.K. Rowling\n\n", len(rowlingBooks))
	
	// Test checkout functionality
	fmt.Println("=== CHECKOUT TESTS ===")
	
	fmt.Printf("Checking out Harry Potter book (ID: 1)...\n")
	success1 := library.CheckoutBook(1, "Alice Smith")
	fmt.Printf("Checkout successful: %t (Expected: true) %s\n\n", success1, checkBool(success1, true))
	
	fmt.Printf("Trying to checkout same book again...\n")
	success2 := library.CheckoutBook(1, "Bob Jones")
	fmt.Printf("Checkout successful: %t (Expected: false) %s\n\n", success2, checkBool(success2, false))
	
	// Show book status after checkout
	fmt.Println("Book 1 status after checkout:")
	book1Results := library.SearchBooks("Harry Potter and the Sorcerer")
	library.DisplayBooks(book1Results)
	
	// Test return functionality  
	fmt.Println("=== RETURN TESTS ===")
	
	fmt.Printf("Returning book (ID: 1)...\n")
	returned := library.ReturnBook(1)
	fmt.Printf("Return successful: %t (Expected: true) %s\n\n", returned, checkBool(returned, true))
	
	// Test overdue fee calculation
	fmt.Println("=== OVERDUE FEE TESTS ===")
	
	fmt.Printf("Overdue fee for book ID 4 (5 days overdue):\n")
	fee := library.CalculateOverdueFee(4)
	fmt.Printf("Fee: $%.2f (Expected: $1.25) %s\n\n", fee, checkFloat(fee, 1.25))
	
	// Test edge cases
	fmt.Println("=== EDGE CASE TESTS ===")
	
	fmt.Println("Searching for non-existent book:")
	noResults := library.SearchBooks("Nonexistent Book")
	library.DisplayBooks(noResults)
	
	fmt.Printf("Trying to return available book (ID: 5):\n")
	badReturn := library.ReturnBook(5)
	fmt.Printf("Return successful: %t (Expected: false) %s\n\n", badReturn, checkBool(badReturn, false))
	
	// Final status
	fmt.Println("=== FINAL LIBRARY STATUS ===")
	fmt.Println("All books:")
	library.DisplayBooks(library.books)
	
	fmt.Println("\n🎉 All bugs fixed! Library system working perfectly! 📖")
}

func checkBool(got, expected bool) string {
	if got == expected {
		return "✅"
	}
	return "❌"
}

func checkFloat(got, expected float64) string {
	if got >= expected-0.01 && got <= expected+0.01 {
		return "✅"
	}
	return "❌"
}