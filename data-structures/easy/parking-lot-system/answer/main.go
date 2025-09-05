package main

import (
	"fmt"
	"sort"
	"time"
)

type Vehicle struct {
	ID       string
	ParkTime time.Time
}

type ParkingLot struct {
	capacity      int
	occupiedSpots map[int]*Vehicle // spot number -> vehicle
	vehicleToSpot map[string]int   // vehicle ID -> spot number
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity:      capacity,
		occupiedSpots: make(map[int]*Vehicle),
		vehicleToSpot: make(map[string]int),
	}
}

func (p *ParkingLot) Park(vehicleId string) int {
	// Check if vehicle is already parked
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

func (p *ParkingLot) Leave(spotNumber int) bool {
	if spotNumber < 1 || spotNumber > p.capacity {
		return false // Invalid spot number
	}
	
	vehicle, occupied := p.occupiedSpots[spotNumber]
	if !occupied {
		return false // Spot is already empty
	}
	
	delete(p.occupiedSpots, spotNumber)
	delete(p.vehicleToSpot, vehicle.ID)
	return true
}

func (p *ParkingLot) GetNearestAvailable() int {
	for i := 1; i <= p.capacity; i++ {
		if _, occupied := p.occupiedSpots[i]; !occupied {
			return i
		}
	}
	return 0 // No available spots
}

func (p *ParkingLot) IsSpotAvailable(spotNumber int) bool {
	if spotNumber < 1 || spotNumber > p.capacity {
		return false // Invalid spot number
	}
	_, occupied := p.occupiedSpots[spotNumber]
	return !occupied
}

func (p *ParkingLot) GetOccupiedSpots() []int {
	spots := make([]int, 0, len(p.occupiedSpots))
	for spot := range p.occupiedSpots {
		spots = append(spots, spot)
	}
	sort.Ints(spots) // Return in sorted order
	return spots
}

func (p *ParkingLot) GetTotalCapacity() int {
	return p.capacity
}

func (p *ParkingLot) GetAvailableCount() int {
	return p.capacity - len(p.occupiedSpots)
}

func (p *ParkingLot) GetVehicleSpot(vehicleId string) int {
	spot, exists := p.vehicleToSpot[vehicleId]
	if !exists {
		return 0 // Vehicle not found
	}
	return spot
}

func (p *ParkingLot) IsFull() bool {
	return len(p.occupiedSpots) == p.capacity
}

func (p *ParkingLot) IsEmpty() bool {
	return len(p.occupiedSpots) == 0
}

func main() {
	fmt.Println("🚗 Parking Lot System Results 🚗\n")

	parking := NewParkingLot(5)
	fmt.Printf("Total capacity: %d", parking.GetTotalCapacity())
	if parking.GetTotalCapacity() == 5 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Available spots: %d", parking.GetAvailableCount())
	if parking.GetAvailableCount() == 5 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Is empty: %t", parking.IsEmpty())
	if parking.IsEmpty() {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Is full: %t", parking.IsFull())
	if !parking.IsFull() {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Park some vehicles
	fmt.Println("\nParking vehicles:")
	spot1 := parking.Park("CAR001")
	fmt.Printf("CAR001 parked in spot: %d", spot1)
	if spot1 == 1 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	spot2 := parking.Park("CAR002")
	fmt.Printf("CAR002 parked in spot: %d", spot2)
	if spot2 == 2 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	spot3 := parking.Park("CAR003")
	fmt.Printf("CAR003 parked in spot: %d", spot3)
	if spot3 == 3 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	fmt.Printf("\nAfter parking 3 cars:\n")
	fmt.Printf("Available spots: %d", parking.GetAvailableCount())
	if parking.GetAvailableCount() == 2 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Is empty: %t", parking.IsEmpty())
	if !parking.IsEmpty() {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Is full: %t", parking.IsFull())
	if !parking.IsFull() {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test spot availability
	fmt.Printf("\nSpot availability:\n")
	fmt.Printf("Spot 1 available: %t", parking.IsSpotAvailable(1))
	if !parking.IsSpotAvailable(1) {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Spot 4 available: %t", parking.IsSpotAvailable(4))
	if parking.IsSpotAvailable(4) {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Find vehicles
	fmt.Printf("\nFinding vehicles:\n")
	fmt.Printf("CAR001 is in spot: %d", parking.GetVehicleSpot("CAR001"))
	if parking.GetVehicleSpot("CAR001") == 1 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("CAR999 is in spot: %d", parking.GetVehicleSpot("CAR999"))
	if parking.GetVehicleSpot("CAR999") == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Get occupied spots
	occupied := parking.GetOccupiedSpots()
	fmt.Printf("\nOccupied spots: %v", occupied)
	expectedOccupied := []int{1, 2, 3}
	if len(occupied) == len(expectedOccupied) {
		match := true
		for i, spot := range occupied {
			if spot != expectedOccupied[i] {
				match = false
				break
			}
		}
		if match {
			fmt.Println(" ✅")
		} else {
			fmt.Println(" ❌")
		}
	} else {
		fmt.Println(" ❌")
	}

	// Vehicle leaves
	fmt.Printf("\nCAR002 leaving spot 2: %t", parking.Leave(2))
	if parking.Leave(2) == false { // Already left
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Available spots after CAR002 leaves: %d", parking.GetAvailableCount())
	if parking.GetAvailableCount() == 3 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Nearest available spot: %d", parking.GetNearestAvailable())
	if parking.GetNearestAvailable() == 2 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test edge cases
	fmt.Printf("\nTesting edge cases:\n")
	fmt.Printf("Leave from invalid spot (0): %t", parking.Leave(0))
	if parking.Leave(0) == false {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Leave from invalid spot (10): %t", parking.Leave(10))
	if parking.Leave(10) == false {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Is spot 0 available: %t", parking.IsSpotAvailable(0))
	if parking.IsSpotAvailable(0) == false {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Printf("Is spot 10 available: %t", parking.IsSpotAvailable(10))
	if parking.IsSpotAvailable(10) == false {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	// Test parking same vehicle twice
	fmt.Printf("Try to park CAR001 again: %d", parking.Park("CAR001"))
	if parking.Park("CAR001") == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}

	fmt.Println("\nAll tests completed! 🅿️")
}