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

// FIXED: All bugs have been corrected!
func CalculateTotal(orders []Order) float64 {
	var subtotal float64
	
	// Calculate subtotal - FIXED: Now properly accumulating
	for _, order := range orders {
		subtotal += order.Price * float64(order.Quantity) // FIXED: Using += instead of =
	}
	
	// Apply discount for orders over $20 - FIXED: Changed >= to >
	discountedTotal := subtotal
	if subtotal > 20 { // FIXED: Now correctly checks > 20, not >= 20
		discount := subtotal * 0.1
		discountedTotal = subtotal - discount // FIXED: Subtracting discount, not adding
	}
	
	// Add sales tax (8%) - FIXED: Now adding tax instead of subtracting
	taxAmount := discountedTotal * 0.08
	finalTotal := discountedTotal + taxAmount // FIXED: Adding tax, not subtracting
	
	// Round to 2 decimal places - FIXED: Using 100 for 2 decimal places
	return math.Round(finalTotal*100) / 100 // FIXED: 100 instead of 10 for 2 decimals
}

func main() {
	fmt.Println("☕ Coffee Shop Calculator (FIXED VERSION) ☕\n")
	
	// Test Case 1: 3 Lattes ($4.50 each)
	// Subtotal: 13.50, No discount, Tax: 1.08, Total: 14.58
	order1 := []Order{
		{ItemName: "Latte", Price: 4.50, Quantity: 3},
	}
	total1 := CalculateTotal(order1)
	fmt.Printf("Order 1 (3 Lattes @ $4.50):\n")
	fmt.Printf("  Subtotal: $13.50, No discount, Tax: $1.08\n")
	fmt.Printf("  Your calculation: $%.2f\n", total1)
	fmt.Printf("  Expected: $14.58\n")
	fmt.Printf("  Status: %s\n\n", checkResult(total1, 14.58))
	
	// Test Case 2: 5 Cappuccinos ($5.00 each)
	// Subtotal: 25.00, Discount: 2.50, After discount: 22.50, Tax: 1.80, Total: 24.30
	order2 := []Order{
		{ItemName: "Cappuccino", Price: 5.00, Quantity: 5},
	}
	total2 := CalculateTotal(order2)
	fmt.Printf("Order 2 (5 Cappuccinos @ $5.00):\n")
	fmt.Printf("  Subtotal: $25.00, Discount: $2.50, After discount: $22.50, Tax: $1.80\n")
	fmt.Printf("  Your calculation: $%.2f\n", total2)
	fmt.Printf("  Expected: $24.30\n")
	fmt.Printf("  Status: %s\n\n", checkResult(total2, 24.30))
	
	// Test Case 3: 1 Espresso ($2.00)
	// Subtotal: 2.00, No discount, Tax: 0.16, Total: 2.16
	order3 := []Order{
		{ItemName: "Espresso", Price: 2.00, Quantity: 1},
	}
	total3 := CalculateTotal(order3)
	fmt.Printf("Order 3 (1 Espresso @ $2.00):\n")
	fmt.Printf("  Subtotal: $2.00, No discount, Tax: $0.16\n")
	fmt.Printf("  Your calculation: $%.2f\n", total3)
	fmt.Printf("  Expected: $2.16\n")
	fmt.Printf("  Status: %s\n\n", checkResult(total3, 2.16))
	
	// Test Case 4: Multiple items ($15.00 total)
	// Subtotal: 15.00, No discount, Tax: 1.20, Total: 16.20
	order4 := []Order{
		{ItemName: "Latte", Price: 4.50, Quantity: 2},
		{ItemName: "Muffin", Price: 3.00, Quantity: 2},
	}
	total4 := CalculateTotal(order4)
	fmt.Printf("Order 4 (2 Lattes @ $4.50, 2 Muffins @ $3.00):\n")
	fmt.Printf("  Subtotal: $15.00, No discount, Tax: $1.20\n")
	fmt.Printf("  Your calculation: $%.2f\n", total4)
	fmt.Printf("  Expected: $16.20\n")
	fmt.Printf("  Status: %s\n\n", checkResult(total4, 16.20))
	
	// Test Case 5: Exactly $20 (edge case - should NOT get discount)
	// Subtotal: 20.00, No discount (not > 20), Tax: 1.60, Total: 21.60
	order5 := []Order{
		{ItemName: "Cappuccino", Price: 5.00, Quantity: 4},
	}
	total5 := CalculateTotal(order5)
	fmt.Printf("Order 5 (4 Cappuccinos @ $5.00 - Edge case):\n")
	fmt.Printf("  Subtotal: $20.00, No discount (not > $20), Tax: $1.60\n")
	fmt.Printf("  Your calculation: $%.2f\n", total5)
	fmt.Printf("  Expected: $21.60\n")
	fmt.Printf("  Status: %s\n", checkResult(total5, 21.60))
	
	fmt.Printf("\n🎉 All bugs fixed! Coffee shop is back in business! ☕\n")
}

func checkResult(got, expected float64) string {
	if math.Abs(got-expected) < 0.01 {
		return "✅ PASS"
	}
	return "❌ FAIL"
}