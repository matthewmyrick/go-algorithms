# Restaurant Queue Manager - Solution 🍽️👥

## Approach

We implement a **queue** using a slice to store customers in FIFO (First In, First Out) order. This is the most straightforward approach for a queue in Go.

## Data Structure Design

```go
type Customer struct {
    Name      string
    PartySize int
    JoinTime  time.Time
}

type RestaurantQueue struct {
    customers []*Customer
}
```

## Key Operations

1. **Enqueue (AddCustomer)**: Append to end of slice
2. **Dequeue (SeatNextCustomer)**: Remove from front of slice
3. **Peek**: Return first element without removing
4. **Size/Empty checks**: Use slice length

## Algorithm Details

### AddCustomer(name, partySize)
- Create new Customer with current timestamp
- Append to customers slice
- Time: O(1) amortized

### SeatNextCustomer()
- Check if queue is empty
- Return first customer and remove from slice
- Time: O(n) due to slice re-indexing

### PeekNext()
- Return first customer without modifying slice
- Time: O(1)

**Time Complexity:**
- AddCustomer: O(1) amortized
- SeatNextCustomer: O(n) 
- PeekNext: O(1)
- GetQueueLength: O(1)
- IsEmpty: O(1)

**Space Complexity:** O(n) where n is number of customers

## Performance Optimization

For high-performance scenarios, consider:
1. **Ring buffer**: Fixed-size circular buffer
2. **Linked list**: O(1) dequeue operations
3. **Two slices**: One for adding, one for removing

For this restaurant scenario, the slice approach is simple and sufficient!