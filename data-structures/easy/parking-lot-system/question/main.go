package main

import (
	"fmt"
	"time"
)

type Vehicle struct {
	ID       string
	ParkTime time.Time
}

type ParkingLot struct {
	// TODO: Define your fields here
}

func NewParkingLot(capacity int) *ParkingLot {
	// TODO: Initialize parking lot with given capacity
	return nil
}

func (p *ParkingLot) Park(vehicleId string) int {
	// TODO: Park vehicle, return spot number (0 if no space)
	return 0
}

func (p *ParkingLot) Leave(spotNumber int) bool {
	// TODO: Remove vehicle from spot, return true if successful
	return false
}

func (p *ParkingLot) GetNearestAvailable() int {
	// TODO: Return nearest available spot number (0 if none)
	return 0
}

func (p *ParkingLot) IsSpotAvailable(spotNumber int) bool {
	// TODO: Check if specific spot is available
	return false
}

func (p *ParkingLot) GetOccupiedSpots() []int {
	// TODO: Return slice of occupied spot numbers
	return nil
}

func (p *ParkingLot) GetTotalCapacity() int {
	// TODO: Return total parking capacity
	return 0
}

func (p *ParkingLot) GetAvailableCount() int {
	// TODO: Return number of available spots
	return 0
}

func (p *ParkingLot) GetVehicleSpot(vehicleId string) int {
	// TODO: Return spot number for vehicle (0 if not found)
	return 0
}

func (p *ParkingLot) IsFull() bool {
	// TODO: Check if parking lot is full
	return false
}

func (p *ParkingLot) IsEmpty() bool {
	// TODO: Check if parking lot is empty
	return true
}

func main() {
	// Test your implementation
	fmt.Println("🚗 Testing Parking Lot System 🚗\n")

	parking := NewParkingLot(5)
	fmt.Printf("Total capacity: %d\n", parking.GetTotalCapacity())
	fmt.Printf("Available spots: %d\n", parking.GetAvailableCount())
	fmt.Printf("Is empty: %t\n", parking.IsEmpty())
	fmt.Printf("Is full: %t\n", parking.IsFull())

	// Park some vehicles
	fmt.Println("\nParking vehicles:")
	spot1 := parking.Park("CAR001")
	fmt.Printf("CAR001 parked in spot: %d\n", spot1)
	
	spot2 := parking.Park("CAR002")
	fmt.Printf("CAR002 parked in spot: %d\n", spot2)
	
	spot3 := parking.Park("CAR003")
	fmt.Printf("CAR003 parked in spot: %d\n", spot3)

	fmt.Printf("\nAfter parking 3 cars:\n")
	fmt.Printf("Available spots: %d\n", parking.GetAvailableCount())
	fmt.Printf("Is empty: %t\n", parking.IsEmpty())
	fmt.Printf("Is full: %t\n", parking.IsFull())

	// Test spot availability
	fmt.Printf("\nSpot availability:\n")
	fmt.Printf("Spot 1 available: %t\n", parking.IsSpotAvailable(1))
	fmt.Printf("Spot 4 available: %t\n", parking.IsSpotAvailable(4))

	// Find vehicles
	fmt.Printf("\nFinding vehicles:\n")
	fmt.Printf("CAR001 is in spot: %d\n", parking.GetVehicleSpot("CAR001"))
	fmt.Printf("CAR999 is in spot: %d\n", parking.GetVehicleSpot("CAR999"))

	// Get occupied spots
	fmt.Printf("\nOccupied spots: %v\n", parking.GetOccupiedSpots())

	// Vehicle leaves
	fmt.Printf("\nCAR002 leaving spot 2: %t\n", parking.Leave(2))
	fmt.Printf("Available spots after CAR002 leaves: %d\n", parking.GetAvailableCount())
	fmt.Printf("Nearest available spot: %d\n", parking.GetNearestAvailable())

	// Park new vehicle (should take spot 2)
	spot4 := parking.Park("CAR004")
	fmt.Printf("CAR004 parked in spot: %d\n", spot4)

	// Fill up the parking lot
	fmt.Println("\nFilling up the parking lot:")
	parking.Park("CAR005")
	parking.Park("CAR006")
	fmt.Printf("Is full: %t\n", parking.IsFull())
	fmt.Printf("Available spots: %d\n", parking.GetAvailableCount())

	// Try to park when full
	fmt.Printf("Trying to park when full: %d\n", parking.Park("CAR007"))

	// Test edge cases
	fmt.Printf("\nTesting edge cases:\n")
	fmt.Printf("Leave from invalid spot (0): %t\n", parking.Leave(0))
	fmt.Printf("Leave from invalid spot (10): %t\n", parking.Leave(10))
	fmt.Printf("Is spot 0 available: %t\n", parking.IsSpotAvailable(0))
	fmt.Printf("Is spot 10 available: %t\n", parking.IsSpotAvailable(10))

	fmt.Printf("\nFinal occupied spots: %v\n", parking.GetOccupiedSpots())
}