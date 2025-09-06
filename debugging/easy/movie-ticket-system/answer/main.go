package main

import (
	"fmt"
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

// FIXED: All booking logic bugs corrected
func (t *Theater) BookSeat(seatID string) bool {
	row, col, valid := t.parseSeatID(seatID)
	if !valid { // FIXED: Changed to !valid
		return false // Invalid seat ID
	}
	
	if t.seats[row][col] { // FIXED: Simplified boolean check
		return false // Seat already booked
	}
	
	t.seats[row][col] = true // FIXED: Set to true (booked)
	return true
}

// FIXED: All parsing bugs corrected
func (t *Theater) parseSeatID(seatID string) (int, int, bool) {
	if len(seatID) < 2 {
		return 0, 0, false
	}
	
	rowChar := seatID[0]
	colStr := seatID[1:]
	
	// Convert row letter to index (A=0, B=1, etc.)
	row := int(rowChar - 'A')
	if row < 0 || row >= t.rows { // FIXED: Changed && to ||
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
	col = col - 1 // FIXED: Subtract 1 to convert 1-based to 0-based
	
	if col < 0 || col >= t.cols {
		return 0, 0, false
	}
	
	return row, col, true
}

// FIXED: Availability checking corrected
func (t *Theater) IsAvailable(seatID string) bool {
	row, col, valid := t.parseSeatID(seatID)
	if !valid {
		return false
	}
	
	return !t.seats[row][col] // FIXED: Return !booked (available)
}

// FIXED: Cancellation logic corrected  
func (t *Theater) CancelSeat(seatID string) bool {
	row, col, valid := t.parseSeatID(seatID)
	if !valid {
		return false
	}
	
	if !t.seats[row][col] { // Seat not booked
		return false
	}
	
	t.seats[row][col] = false // FIXED: Set to false (available)
	return true
}

// FIXED: Display symbols corrected
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
				fmt.Print("X ") // FIXED: Show "X" for booked (true)
			} else {
				fmt.Print("O ") // FIXED: Show "O" for available (false)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	fmt.Println("🎬 Movie Ticket System (FIXED VERSION) 🎬\n")
	
	// Create a 5x5 theater
	theater := NewTheater(5, 5)
	
	fmt.Println("Initial seating chart:")
	theater.ShowSeatingChart()
	
	// Test seat booking
	fmt.Println("Testing seat bookings:")
	
	// Should succeed
	result1 := theater.BookSeat("A1")
	fmt.Printf("Book A1: %t (Expected: true) %s\n", result1, checkResult(result1, true))
	
	result2 := theater.BookSeat("C3")
	fmt.Printf("Book C3: %t (Expected: true) %s\n", result2, checkResult(result2, true))
	
	// Should fail (already booked)
	result3 := theater.BookSeat("A1")
	fmt.Printf("Book A1 again: %t (Expected: false) %s\n", result3, checkResult(result3, false))
	
	// Should fail (invalid seat)
	result4 := theater.BookSeat("Z9")
	fmt.Printf("Book Z9 (invalid): %t (Expected: false) %s\n", result4, checkResult(result4, false))
	
	fmt.Println("\nSeating chart after bookings:")
	theater.ShowSeatingChart()
	
	// Test availability checking
	fmt.Println("Testing seat availability:")
	avail1 := theater.IsAvailable("A1")
	fmt.Printf("A1 available: %t (Expected: false) %s\n", avail1, checkResult(avail1, false))
	
	avail2 := theater.IsAvailable("A2")
	fmt.Printf("A2 available: %t (Expected: true) %s\n", avail2, checkResult(avail2, true))
	
	avail3 := theater.IsAvailable("C3")
	fmt.Printf("C3 available: %t (Expected: false) %s\n", avail3, checkResult(avail3, false))
	
	// Test cancellation
	fmt.Println("\nTesting cancellations:")
	cancel1 := theater.CancelSeat("A1")
	fmt.Printf("Cancel A1: %t (Expected: true) %s\n", cancel1, checkResult(cancel1, true))
	
	cancel2 := theater.CancelSeat("A2") // Not booked
	fmt.Printf("Cancel A2 (not booked): %t (Expected: false) %s\n", cancel2, checkResult(cancel2, false))
	
	fmt.Println("\nFinal seating chart:")
	theater.ShowSeatingChart()
	
	// Final verification
	fmt.Println("Final verification:")
	a1Final := theater.IsAvailable("A1")
	c3Final := theater.IsAvailable("C3")
	fmt.Printf("A1 available: %t (Expected: true) %s\n", a1Final, checkResult(a1Final, true))
	fmt.Printf("C3 available: %t (Expected: false) %s\n", c3Final, checkResult(c3Final, false))
	
	fmt.Println("\n🎉 All bugs fixed! Movie theater booking system working perfectly! 🍿")
}

func checkResult(got, expected bool) string {
	if got == expected {
		return "✅"
	}
	return "❌"
}