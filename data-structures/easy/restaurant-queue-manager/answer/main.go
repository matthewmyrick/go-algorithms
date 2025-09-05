package main

import (
	"fmt"
	"time"
)

type Customer struct {
	Name      string
	PartySize int
	JoinTime  time.Time
}

type RestaurantQueue struct {
	customers []*Customer
}

func NewRestaurantQueue() *RestaurantQueue {
	return &RestaurantQueue{
		customers: make([]*Customer, 0),
	}
}

func (q *RestaurantQueue) AddCustomer(name string, partySize int) {
	customer := &Customer{
		Name:      name,
		PartySize: partySize,
		JoinTime:  time.Now(),
	}
	q.customers = append(q.customers, customer)
}

func (q *RestaurantQueue) SeatNextCustomer() *Customer {
	if len(q.customers) == 0 {
		return nil
	}
	
	customer := q.customers[0]
	q.customers = q.customers[1:] // Remove first customer
	return customer
}

func (q *RestaurantQueue) PeekNext() *Customer {
	if len(q.customers) == 0 {
		return nil
	}
	return q.customers[0]
}

func (q *RestaurantQueue) GetQueueLength() int {
	return len(q.customers)
}

func (q *RestaurantQueue) IsEmpty() bool {
	return len(q.customers) == 0
}

func (q *RestaurantQueue) GetAllWaiting() []*Customer {
	// Return a copy to prevent external modifications
	result := make([]*Customer, len(q.customers))
	copy(result, q.customers)
	return result
}

func main() {
	fmt.Println("🍽️ Restaurant Queue Manager Results 🍽️\n")

	queue := NewRestaurantQueue()
	fmt.Printf("Is empty: %t", queue.IsEmpty())
	if queue.IsEmpty() {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Queue length: %d", queue.GetQueueLength())
	if queue.GetQueueLength() == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Add customers
	queue.AddCustomer("Alice", 2)
	queue.AddCustomer("Bob", 4)
	queue.AddCustomer("Charlie", 1)
	
	fmt.Printf("\nAfter adding 3 customers:\n")
	fmt.Printf("Queue length: %d", queue.GetQueueLength())
	if queue.GetQueueLength() == 3 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Is empty: %t", queue.IsEmpty())
	if !queue.IsEmpty() {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Peek at next customer
	next := queue.PeekNext()
	if next != nil {
		fmt.Printf("Next customer: %s (party of %d)", next.Name, next.PartySize)
		if next.Name == "Alice" && next.PartySize == 2 {
			fmt.Println(" ✅")
		} else {
			fmt.Println(" ❌")
		}
	}
	
	fmt.Printf("Queue length after peek: %d", queue.GetQueueLength())
	if queue.GetQueueLength() == 3 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Seat customers
	fmt.Println("\nSeating customers:")
	expectedNames := []string{"Alice", "Bob", "Charlie"}
	expectedSizes := []int{2, 4, 1}
	
	for i := 0; !queue.IsEmpty() && i < len(expectedNames); i++ {
		customer := queue.SeatNextCustomer()
		if customer != nil {
			fmt.Printf("Seated: %s (party of %d)", customer.Name, customer.PartySize)
			if customer.Name == expectedNames[i] && customer.PartySize == expectedSizes[i] {
				fmt.Println(" ✅")
			} else {
				fmt.Println(" ❌")
			}
			
			expectedRemaining := len(expectedNames) - i - 1
			fmt.Printf("Remaining in queue: %d", queue.GetQueueLength())
			if queue.GetQueueLength() == expectedRemaining {
				fmt.Println(" ✅")
			} else {
				fmt.Println(" ❌")
			}
		}
	}

	// Test empty queue
	fmt.Printf("\nTesting empty queue:\n")
	fmt.Printf("Peek next: %v", queue.PeekNext())
	if queue.PeekNext() == nil {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Seat next: %v", queue.SeatNextCustomer())
	if queue.SeatNextCustomer() == nil {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test GetAllWaiting
	fmt.Printf("\nTesting GetAllWaiting on empty queue: %d customers", len(queue.GetAllWaiting()))
	if len(queue.GetAllWaiting()) == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
}