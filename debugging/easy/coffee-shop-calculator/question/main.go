package main

import (
	"fmt"
	"math"
)

type Order struct {
	ItemName string
	Price    float64
	Quantity int
}

// BUG: This function has multiple bugs that need to be fixed!
func CalculateTotal(orders []Order) float64 {
	var subtotal float64
	
	// Calculate subtotal
	for _, order := range orders {
		subtotal = order.Price * float64(order.Quantity) // BUG #1: Not accumulating
	}
	
	// Apply discount for orders over $20
	discountedTotal := subtotal
	if subtotal >= 20 { // BUG #2: Should be > not >=
		discount := subtotal * 0.1
		discountedTotal = subtotal + discount // BUG #3: Should subtract, not add
	}
	
	// Add sales tax (8%)
	taxAmount := discountedTotal * 0.08
	finalTotal := discountedTotal - taxAmount // BUG #4: Should add tax, not subtract
	
	// Round to 2 decimal places
	return math.Round(finalTotal*10) / 10 // BUG #5: Should be *100 / 100 for 2 decimal places
}

func main() {
	fmt.Println("☕ Coffee Shop Calculator (BUGGY VERSION) ☕\n")
	
	// Test Case 1: 3 Lattes
	order1 := []Order{
		{ItemName: "Latte", Price: 4.50, Quantity: 3},
	}
	total1 := CalculateTotal(order1)
	fmt.Printf("Order 1 (3 Lattes @ $4.50):\n")
	fmt.Printf("  Your calculation: $%.2f\n", total1)
	fmt.Printf("  Expected: $14.58\n")
	fmt.Printf("  Status: %s\n\n", checkResult(total1, 14.58))
	
	// Test Case 2: 5 Cappuccinos (should get discount)
	order2 := []Order{
		{ItemName: "Cappuccino", Price: 5.00, Quantity: 5},
	}
	total2 := CalculateTotal(order2)
	fmt.Printf("Order 2 (5 Cappuccinos @ $5.00):\n")
	fmt.Printf("  Your calculation: $%.2f\n", total2)
	fmt.Printf("  Expected: $24.30\n")
	fmt.Printf("  Status: %s\n\n", checkResult(total2, 24.30))
	
	// Test Case 3: 1 Espresso
	order3 := []Order{
		{ItemName: "Espresso", Price: 2.00, Quantity: 1},
	}
	total3 := CalculateTotal(order3)
	fmt.Printf("Order 3 (1 Espresso @ $2.00):\n")
	fmt.Printf("  Your calculation: $%.2f\n", total3)
	fmt.Printf("  Expected: $2.16\n")
	fmt.Printf("  Status: %s\n\n", checkResult(total3, 2.16))
	
	// Test Case 4: Multiple items
	order4 := []Order{
		{ItemName: "Latte", Price: 4.50, Quantity: 2},
		{ItemName: "Muffin", Price: 3.00, Quantity: 2},
	}
	total4 := CalculateTotal(order4)
	fmt.Printf("Order 4 (2 Lattes @ $4.50, 2 Muffins @ $3.00):\n")
	fmt.Printf("  Your calculation: $%.2f\n", total4)
	fmt.Printf("  Expected: $16.20\n")
	fmt.Printf("  Status: %s\n\n", checkResult(total4, 16.20))
	
	// Test Case 5: Exactly $20 (edge case)
	order5 := []Order{
		{ItemName: "Cappuccino", Price: 5.00, Quantity: 4},
	}
	total5 := CalculateTotal(order5)
	fmt.Printf("Order 5 (4 Cappuccinos @ $5.00):\n")
	fmt.Printf("  Your calculation: $%.2f\n", total5)
	fmt.Printf("  Expected: $21.60\n")
	fmt.Printf("  Status: %s\n", checkResult(total5, 21.60))
}

func checkResult(got, expected float64) string {
	if math.Abs(got-expected) < 0.01 {
		return "✅ PASS"
	}
	return "❌ FAIL"
}