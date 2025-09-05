# Parking Lot System 🚗🅿️

You're managing a smart parking lot system! The parking lot has a fixed number of spaces, and you need to efficiently track which spaces are occupied, find available spots, and manage vehicle check-ins and check-outs.

## The Problem

Implement a `ParkingLot` system that manages parking spaces with the following operations:

1. **Park(vehicleId)** - Park a vehicle in the nearest available spot (return spot number)
2. **Leave(spotNumber)** - Remove vehicle from specific spot
3. **GetNearestAvailable()** - Get the nearest available parking spot number
4. **IsSpotAvailable(spotNumber)** - Check if specific spot is available
5. **GetOccupiedSpots()** - Get list of all occupied spot numbers
6. **GetTotalCapacity()** - Get total parking capacity
7. **GetAvailableCount()** - Get number of available spots
8. **GetVehicleSpot(vehicleId)** - Find which spot a vehicle is parked in
9. **IsFull()** - Check if parking lot is completely full
10. **IsEmpty()** - Check if parking lot is completely empty

## Vehicle Structure

```go
type Vehicle struct {
    ID       string
    ParkTime time.Time
}
```

## Example Usage

```go
parking := NewParkingLot(5)  // 5 parking spots (1-5)
spot1 := parking.Park("CAR001")  // Returns 1 (first available)
spot2 := parking.Park("CAR002")  // Returns 2 
parking.Leave(1)                 // Spot 1 is now available
spot3 := parking.Park("CAR003")  // Returns 1 (nearest available)

fmt.Println(parking.GetAvailableCount())    // 2
fmt.Println(parking.GetVehicleSpot("CAR002"))  // 2
fmt.Println(parking.IsFull())               // false
```

## Your Task

Implement the ParkingLot struct and all its methods.

**Structure:**
```go
type Vehicle struct {
    ID       string
    ParkTime time.Time
}

type ParkingLot struct {
    // TODO: Define your fields here
}

func NewParkingLot(capacity int) *ParkingLot {
    // TODO: Initialize parking lot with given capacity
}

func (p *ParkingLot) Park(vehicleId string) int {
    // TODO: Park vehicle, return spot number (0 if no space)
}

func (p *ParkingLot) Leave(spotNumber int) bool {
    // TODO: Remove vehicle from spot, return true if successful
}

func (p *ParkingLot) GetNearestAvailable() int {
    // TODO: Return nearest available spot number (0 if none)
}

func (p *ParkingLot) IsSpotAvailable(spotNumber int) bool {
    // TODO: Check if specific spot is available
}

func (p *ParkingLot) GetOccupiedSpots() []int {
    // TODO: Return slice of occupied spot numbers
}

func (p *ParkingLot) GetTotalCapacity() int {
    // TODO: Return total parking capacity
}

func (p *ParkingLot) GetAvailableCount() int {
    // TODO: Return number of available spots
}

func (p *ParkingLot) GetVehicleSpot(vehicleId string) int {
    // TODO: Return spot number for vehicle (0 if not found)
}

func (p *ParkingLot) IsFull() bool {
    // TODO: Check if parking lot is full
}

func (p *ParkingLot) IsEmpty() bool {
    // TODO: Check if parking lot is empty
}
```

## Test Cases

The main function will test various parking operations including:
- Parking vehicles and getting assigned spots
- Finding nearest available spots
- Vehicle check-out and check-in
- Capacity management and edge cases

Drive safely! 🚙💨