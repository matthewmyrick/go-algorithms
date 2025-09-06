package main

import (
	"fmt"
	"strings"
)

type Theater struct {
	rows    int
	cols    int
	seats   [][]bool // true = booked, false = available
}

func NewTheater(rows, cols int) *Theater {
	seats := make([][]bool, rows)
	for i := range seats {
		seats[i] = make([]bool, cols)
	}
	return &Theater{
		rows:  rows,
		cols:  cols,
		seats: seats,
	}
}

// BUG: This function has bugs in seat parsing and booking logic
func (t *Theater) BookSeat(seatID string) bool {
	row, col, valid := t.parseSeatID(seatID)
	if valid { // BUG #1: Should check !valid
		return false // Invalid seat ID
	}
	
	if t.seats[row][col] == true { // BUG #2: Logic is correct but should be clearer
		return false // Seat already booked
	}
	
	t.seats[row][col] = false // BUG #3: Should set to true (booked)
	return true
}

// BUG: This function has parsing bugs
func (t *Theater) parseSeatID(seatID string) (int, int, bool) {
	if len(seatID) < 2 {
		return 0, 0, false
	}
	
	rowChar := seatID[0]
	colStr := seatID[1:]
	
	// Convert row letter to index (A=0, B=1, etc.)
	row := int(rowChar - 'A')
	if row < 0 && row >= t.rows { // BUG #4: Should use OR (||), not AND (&&)
		return 0, 0, false
	}
	
	// Parse column number
	col := 0
	for _, char := range colStr {
		if char >= '0' && char <= '9' {
			col = col*10 + int(char-'0')
		} else {
			return 0, 0, false
		}
	}
	col = col + 1 // BUG #5: Should subtract 1 (convert 1-based to 0-based)
	
	if col < 0 || col >= t.cols {
		return 0, 0, false
	}
	
	return row, col, true
}

// BUG: This function has availability checking bug
func (t *Theater) IsAvailable(seatID string) bool {
	row, col, valid := t.parseSeatID(seatID)
	if !valid {
		return false
	}
	
	return t.seats[row][col] == true // BUG #6: Should check == false (available)
}

// BUG: This function has cancellation bug
func (t *Theater) CancelSeat(seatID string) bool {
	row, col, valid := t.parseSeatID(seatID)
	if !valid {
		return false
	}
	
	if t.seats[row][col] == false { // Seat not booked
		return false
	}
	
	t.seats[row][col] = true // BUG #7: Should set to false (available)
	return true
}

// BUG: This function has display bugs
func (t *Theater) ShowSeatingChart() {
	fmt.Println("\nSeating Chart (O=Available, X=Booked):")
	fmt.Print("   ")
	for col := 1; col <= t.cols; col++ {
		fmt.Printf("%d ", col)
	}
	fmt.Println()
	
	for row := 0; row < t.rows; row++ {
		fmt.Printf("%c  ", 'A'+row)
		for col := 0; col < t.cols; col++ {
			if t.seats[row][col] {
				fmt.Print("O ") // BUG #8: Should show "X" for booked
			} else {
				fmt.Print("X ") // BUG #9: Should show "O" for available
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	fmt.Println("🎬 Movie Ticket System (BUGGY VERSION) 🎬\n")
	
	// Create a 5x5 theater
	theater := NewTheater(5, 5)
	
	fmt.Println("Initial seating chart:")
	theater.ShowSeatingChart()
	
	// Test seat booking
	fmt.Println("Testing seat bookings:")
	
	// Should succeed
	result1 := theater.BookSeat("A1")
	fmt.Printf("Book A1: %t (Expected: true)\n", result1)
	
	result2 := theater.BookSeat("C3")
	fmt.Printf("Book C3: %t (Expected: true)\n", result2)
	
	// Should fail (already booked)
	result3 := theater.BookSeat("A1")
	fmt.Printf("Book A1 again: %t (Expected: false)\n", result3)
	
	// Should fail (invalid seat)
	result4 := theater.BookSeat("Z9")
	fmt.Printf("Book Z9 (invalid): %t (Expected: false)\n", result4)
	
	fmt.Println("\\nSeating chart after bookings:")
	theater.ShowSeatingChart()
	
	// Test availability checking
	fmt.Println("Testing seat availability:")
	fmt.Printf("A1 available: %t (Expected: false)\\n", theater.IsAvailable("A1"))
	fmt.Printf("A2 available: %t (Expected: true)\\n", theater.IsAvailable("A2"))
	fmt.Printf("C3 available: %t (Expected: false)\\n", theater.IsAvailable("C3"))
	
	// Test cancellation
	fmt.Println("\\nTesting cancellations:")
	cancel1 := theater.CancelSeat("A1")
	fmt.Printf("Cancel A1: %t (Expected: true)\\n", cancel1)
	
	cancel2 := theater.CancelSeat("A2") // Not booked
	fmt.Printf("Cancel A2 (not booked): %t (Expected: false)\\n", cancel2)
	
	fmt.Println("\\nFinal seating chart:")
	theater.ShowSeatingChart()
	
	fmt.Println("Expected: A1 should be available (O), C3 should be booked (X)")
}