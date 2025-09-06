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
			{4, "To Kill a Mockingbird", "Harper Lee", "978-0061120084", false, "John Doe", time.Now().AddDate(0, 0, -5)}, // 5 days overdue
			{5, "1984", "George Orwell", "978-0452284234", true, "", time.Time{}},
		},
	}
}

// BUG: This function has string searching bugs
func (l *Library) SearchBooks(query string) []Book {
	var results []Book
	query = strings.ToUpper(query) // BUG #1: Should use ToLower for consistent comparison
	
	for _, book := range l.books {
		// Check title, author, and ISBN
		if strings.Contains(strings.ToLower(book.Title), query) && // BUG #2: Mixed case comparison (query is upper)
		   strings.Contains(strings.ToLower(book.Author), query) && // BUG #3: Should use OR (||), not AND (&&)
		   strings.Contains(strings.ToLower(book.ISBN), query) {
			results = append(results, book)
		}
	}
	
	return results
}

// BUG: This function has checkout logic bugs
func (l *Library) CheckoutBook(bookID int, patronName string) bool {
	for i := range l.books {
		if l.books[i].ID == bookID {
			if l.books[i].IsAvailable == false { // BUG #4: Should check if available (true), not unavailable
				return false // Book not available
			}
			
			l.books[i].IsAvailable = true // BUG #5: Should set to false (checked out)
			l.books[i].CheckedOutBy = patronName
			l.books[i].DueDate = time.Now().AddDate(0, 0, 7) // BUG #6: Should be 14 days, not 7
			return true
		}
	}
	return false // Book not found
}

// BUG: This function has return logic bugs
func (l *Library) ReturnBook(bookID int) bool {
	for i := range l.books {
		if l.books[i].ID == bookID {
			if l.books[i].IsAvailable { // BUG #7: Should check if NOT available (!available)
				return false // Book wasn't checked out
			}
			
			l.books[i].IsAvailable = false // BUG #8: Should set to true (available)
			l.books[i].CheckedOutBy = "RETURNED" // BUG #9: Should clear to empty string
			l.books[i].DueDate = time.Now() // BUG #10: Should reset to zero time
			return true
		}
	}
	return false // Book not found
}

// BUG: This function has overdue calculation bugs
func (l *Library) CalculateOverdueFee(bookID int) float64 {
	for _, book := range l.books {
		if book.ID == bookID {
			if book.IsAvailable { // Book not checked out
				return 0.0
			}
			
			if time.Now().Before(book.DueDate) { // Not overdue
				return 0.0
			}
			
			// Calculate overdue days
			overdueDays := time.Now().Sub(book.DueDate).Hours() * 24 // BUG #11: Wrong calculation (hours * 24 instead of hours / 24)
			feePerDay := 0.25
			return overdueDays * feePerDay
		}
	}
	return 0.0 // Book not found
}

// BUG: This function has display formatting bugs
func (l *Library) DisplayBooks(books []Book) {
	if len(books) > 0 { // BUG #12: Should check == 0 for empty results
		fmt.Println("No books found.")
		return
	}
	
	for _, book := range books {
		status := "Available"
		if book.IsAvailable { // BUG #13: Logic is backwards
			status = "Checked out"
		}
		
		fmt.Printf("ID: %d\\n", book.ID)
		fmt.Printf("Title: %s\\n", book.Title)
		fmt.Printf("Author: %s\\n", book.Author)
		fmt.Printf("Status: %s\\n", status)
		
		if book.IsAvailable { // BUG #14: Should check !IsAvailable for checked out books
			fmt.Printf("Checked out by: %s\\n", book.CheckedOutBy)
			fmt.Printf("Due date: %s\\n", book.DueDate.Format("2006-01-02"))
		}
		fmt.Println("---")
	}
}

func main() {
	fmt.Println("📚 Library Book Finder (BUGGY VERSION) 📚\\n")
	
	library := NewLibrary()
	
	// Test search functionality
	fmt.Println("=== SEARCH TESTS ===")
	
	fmt.Println("Searching for 'Harry Potter':")
	harryPotterBooks := library.SearchBooks("Harry Potter")
	library.DisplayBooks(harryPotterBooks)
	fmt.Printf("Expected: 2 Harry Potter books\\n\\n")
	
	fmt.Println("Searching for 'J.K. Rowling':")
	rowlingBooks := library.SearchBooks("J.K. Rowling")
	library.DisplayBooks(rowlingBooks)
	fmt.Printf("Expected: 2 books by J.K. Rowling\\n\\n")
	
	// Test checkout functionality
	fmt.Println("=== CHECKOUT TESTS ===")
	
	fmt.Printf("Checking out Harry Potter book (ID: 1)...\\n")
	success1 := library.CheckoutBook(1, "Alice Smith")
	fmt.Printf("Checkout successful: %t (Expected: true)\\n\\n", success1)
	
	fmt.Printf("Trying to checkout same book again...\\n")
	success2 := library.CheckoutBook(1, "Bob Jones")
	fmt.Printf("Checkout successful: %t (Expected: false)\\n\\n", success2)
	
	// Test return functionality  
	fmt.Println("=== RETURN TESTS ===")
	
	fmt.Printf("Returning book (ID: 1)...\\n")
	returned := library.ReturnBook(1)
	fmt.Printf("Return successful: %t (Expected: true)\\n\\n", returned)
	
	// Test overdue fee calculation
	fmt.Println("=== OVERDUE FEE TESTS ===")
	
	fmt.Printf("Overdue fee for book ID 4 (5 days overdue):\\n")
	fee := library.CalculateOverdueFee(4)
	fmt.Printf("Fee: $%.2f (Expected: $1.25)\\n\\n", fee)
	
	// Final status
	fmt.Println("=== FINAL LIBRARY STATUS ===")
	fmt.Println("All books:")
	library.DisplayBooks(library.books)
}