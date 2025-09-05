# Restaurant Queue Manager 🍽️👥

You're managing the waiting queue at the hottest restaurant in town! Customers arrive and need to be seated in the order they arrived (FIFO - First In, First Out). You need to implement a queue system to manage customer flow efficiently.

## The Problem

Implement a `RestaurantQueue` that manages customers waiting for tables. The system should support:

1. **AddCustomer(name, partySize)** - Add a customer to the queue
2. **SeatNextCustomer()** - Seat the next customer (remove from queue)
3. **PeekNext()** - See who's next without removing them
4. **GetQueueLength()** - Get the current number of waiting customers
5. **IsEmpty()** - Check if the queue is empty
6. **GetAllWaiting()** - Get list of all customers currently waiting

## Customer Structure

```go
type Customer struct {
    Name      string
    PartySize int
    JoinTime  time.Time
}
```

## Example Usage

```go
queue := NewRestaurantQueue()
queue.AddCustomer("Alice", 2)
queue.AddCustomer("Bob", 4)
queue.AddCustomer("Charlie", 1)

fmt.Println(queue.PeekNext())        // {Alice 2 2024-01-01...}
fmt.Println(queue.GetQueueLength())  // 3

seated := queue.SeatNextCustomer()   // {Alice 2 2024-01-01...}
fmt.Println(queue.GetQueueLength())  // 2
```

## Your Task

Implement the RestaurantQueue struct and all its methods.

**Structure:**
```go
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
}

func (q *RestaurantQueue) AddCustomer(name string, partySize int) {
    // TODO: Add customer to queue
}

func (q *RestaurantQueue) SeatNextCustomer() *Customer {
    // TODO: Remove and return next customer
}

func (q *RestaurantQueue) PeekNext() *Customer {
    // TODO: Return next customer without removing
}

func (q *RestaurantQueue) GetQueueLength() int {
    // TODO: Return current queue size
}

func (q *RestaurantQueue) IsEmpty() bool {
    // TODO: Check if queue is empty
}

func (q *RestaurantQueue) GetAllWaiting() []*Customer {
    // TODO: Return all waiting customers
}
```

## Test Cases

The main function will test various queue operations including:
- Adding customers to the queue
- Seating customers in FIFO order
- Handling empty queue scenarios
- Peeking at next customer without removal

Keep those customers happy! 😊🎉