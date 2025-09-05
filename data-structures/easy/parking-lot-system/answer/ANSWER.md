# Parking Lot System - Solution 🚗🅿️

## Approach

We use a **map to track occupied spots** and implement efficient algorithms for finding the nearest available spot. This approach balances space efficiency with fast lookups.

## Data Structure Design

```go
type Vehicle struct {
    ID       string
    ParkTime time.Time
}

type ParkingLot struct {
    capacity        int
    occupiedSpots   map[int]*Vehicle    // spot number -> vehicle
    vehicleToSpot   map[string]int      // vehicle ID -> spot number
}
```

## Key Operations

1. **Park**: Find nearest available spot, assign vehicle
2. **Leave**: Remove vehicle from spot mappings
3. **Nearest Available**: Linear search for lowest unoccupied number
4. **Lookups**: Use maps for O(1) access

## Algorithm Details

### Finding Nearest Available Spot
```go
func (p *ParkingLot) GetNearestAvailable() int {
    for i := 1; i <= p.capacity; i++ {
        if _, occupied := p.occupiedSpots[i]; !occupied {
            return i
        }
    }
    return 0 // No available spots
}
```

### Parking a Vehicle
```go
func (p *ParkingLot) Park(vehicleId string) int {
    // Check if vehicle already parked
    if _, exists := p.vehicleToSpot[vehicleId]; exists {
        return 0 // Vehicle already parked
    }
    
    spot := p.GetNearestAvailable()
    if spot == 0 {
        return 0 // Lot is full
    }
    
    vehicle := &Vehicle{
        ID:       vehicleId,
        ParkTime: time.Now(),
    }
    
    p.occupiedSpots[spot] = vehicle
    p.vehicleToSpot[vehicleId] = spot
    return spot
}
```

### Efficient Occupied Spots Retrieval
```go
func (p *ParkingLot) GetOccupiedSpots() []int {
    spots := make([]int, 0, len(p.occupiedSpots))
    for spot := range p.occupiedSpots {
        spots = append(spots, spot)
    }
    sort.Ints(spots) // Return in sorted order
    return spots
}
```

**Time Complexity:**
- Park: O(n) worst case (finding nearest available)
- Leave: O(1)
- GetNearestAvailable: O(n) worst case
- IsSpotAvailable: O(1)
- GetVehicleSpot: O(1)
- GetOccupiedSpots: O(k log k) where k = occupied spots

**Space Complexity:** O(k) where k is number of occupied spots

## Optimizations

For larger parking lots, consider:
1. **Min-heap**: Track available spots for O(log n) nearest available
2. **Bit vector**: For memory-efficient availability tracking
3. **Segment tree**: For range-based availability queries
4. **Priority queue**: For complex spot assignment rules

## Real-world Considerations

1. **Reservation system**: Pre-book spots
2. **Vehicle types**: Different sizes (compact, SUV, etc.)
3. **Pricing tiers**: Premium spots cost more
4. **Time limits**: Enforce maximum parking duration

The map-based approach is perfect for typical parking lot sizes!