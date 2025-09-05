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
	// TODO: Define your fields here
}

func NewRestaurantQueue() *RestaurantQueue {
	// TODO: Initialize and return new queue
	return nil
}

func (q *RestaurantQueue) AddCustomer(name string, partySize int) {
	// TODO: Add customer to queue
}

func (q *RestaurantQueue) SeatNextCustomer() *Customer {
	// TODO: Remove and return next customer
	return nil
}

func (q *RestaurantQueue) PeekNext() *Customer {
	// TODO: Return next customer without removing
	return nil
}

func (q *RestaurantQueue) GetQueueLength() int {
	// TODO: Return current queue size
	return 0
}

func (q *RestaurantQueue) IsEmpty() bool {
	// TODO: Check if queue is empty
	return true
}

func (q *RestaurantQueue) GetAllWaiting() []*Customer {
	// TODO: Return all waiting customers
	return nil
}

func main() {
	// Test your implementation
	fmt.Println("🍽️ Testing Restaurant Queue Manager 🍽️\n")

	queue := NewRestaurantQueue()
	fmt.Printf("Is empty: %t\n", queue.IsEmpty())
	fmt.Printf("Queue length: %d\n\n", queue.GetQueueLength())

	// Add customers
	queue.AddCustomer("Alice", 2)
	queue.AddCustomer("Bob", 4)
	queue.AddCustomer("Charlie", 1)
	
	fmt.Printf("After adding 3 customers:\n")
	fmt.Printf("Queue length: %d\n", queue.GetQueueLength())
	fmt.Printf("Is empty: %t\n", queue.IsEmpty())

	// Peek at next customer
	if next := queue.PeekNext(); next != nil {
		fmt.Printf("Next customer: %s (party of %d)\n", next.Name, next.PartySize)
	}
	fmt.Printf("Queue length after peek: %d\n\n", queue.GetQueueLength())

	// Seat customers
	fmt.Println("Seating customers:")
	for !queue.IsEmpty() {
		customer := queue.SeatNextCustomer()
		if customer != nil {
			fmt.Printf("Seated: %s (party of %d)\n", customer.Name, customer.PartySize)
			fmt.Printf("Remaining in queue: %d\n", queue.GetQueueLength())
		}
	}

	// Test empty queue
	fmt.Printf("\nTesting empty queue:\n")
	fmt.Printf("Peek next: %v\n", queue.PeekNext())
	fmt.Printf("Seat next: %v\n", queue.SeatNextCustomer())
}