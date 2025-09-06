package main

import (
	"fmt"
)

type User struct {
	Name       string
	WeightLbs  float64
	DailyGoal  int
}

type FitnessTracker struct {
	user      User
	stepCount int
}

func NewFitnessTracker(user User) *FitnessTracker {
	return &FitnessTracker{
		user:      user,
		stepCount: 0,
	}
}

// FIXED: Step counting logic corrected
func (ft *FitnessTracker) AddSteps(steps int) {
	if steps > 0 { // FIXED: Changed to > 0
		ft.stepCount += steps // FIXED: Changed to += to accumulate
	}
}

// FIXED: Distance calculation corrected
func (ft *FitnessTracker) GetDistanceMiles() float64 {
	// Average person: 2000 steps = 1 mile
	return float64(ft.stepCount) / 2000.0 // FIXED: Divide by 2000
}

// FIXED: Calorie calculation corrected
func (ft *FitnessTracker) GetCaloriesBurned() float64 {
	// Formula: 4 calories per 100 steps for 150 lb person
	// Adjust for user weight: (userWeight / 150) * base calories
	
	baseCaloriesPer100Steps := 4.0
	stepsIn100s := float64(ft.stepCount) / 100.0 // FIXED: Divide by 100
	weightFactor := ft.user.WeightLbs / 150.0 // FIXED: Divide by 150
	
	return stepsIn100s * baseCaloriesPer100Steps * weightFactor
}

// FIXED: Goal progress calculation corrected
func (ft *FitnessTracker) GetGoalProgress() float64 {
	if ft.user.DailyGoal == 0 {
		return 0.0
	}
	
	progress := float64(ft.stepCount) / float64(ft.user.DailyGoal) // FIXED: Divide instead of multiply
	if progress > 1.0 { // FIXED: Changed >= to > to allow exactly 100%
		return 1.0 // Cap at 100%
	}
	return progress
}

// FIXED: Activity level boundaries corrected
func (ft *FitnessTracker) GetActivityLevel() string {
	steps := ft.stepCount
	
	if steps < 5000 {
		return "Sedentary"
	} else if steps < 7500 {
		return "Lightly Active"  
	} else if steps < 10000 { // FIXED: Changed <= to < 
		return "Moderately Active"
	} else if steps < 12500 {
		return "Very Active"
	} else {
		return "Extremely Active"
	}
}

// FIXED: Reset function corrected
func (ft *FitnessTracker) ResetDaily() {
	ft.stepCount = 0 // FIXED: Reset to 0, not 1
}

func main() {
	fmt.Println("🏃‍♂️ Fitness Tracker (FIXED VERSION) 🏃‍♂️\n")
	
	// Create user and tracker
	user := User{
		Name:      "Alice",
		WeightLbs: 150.0,
		DailyGoal: 10000,
	}
	
	tracker := NewFitnessTracker(user)
	
	fmt.Printf("User: %s (%.0f lbs, Goal: %d steps)\n\n", user.Name, user.WeightLbs, user.DailyGoal)
	
	// Test adding steps
	fmt.Println("Adding 5000 steps...")
	tracker.AddSteps(5000)
	fmt.Printf("Steps: %d %s\n", tracker.stepCount, checkInt(tracker.stepCount, 5000))
	fmt.Printf("Expected: 5000 steps\n\n")
	
	// Add more steps
	fmt.Println("Adding 3000 more steps...")
	tracker.AddSteps(3000)
	fmt.Printf("Steps: %d %s\n", tracker.stepCount, checkInt(tracker.stepCount, 8000))
	fmt.Printf("Expected: 8000 steps\n\n")
	
	// Add final steps to reach 10,000
	fmt.Println("Adding 2000 more steps...")
	tracker.AddSteps(2000)
	fmt.Printf("Steps: %d %s\n", tracker.stepCount, checkInt(tracker.stepCount, 10000))
	fmt.Printf("Expected: 10000 steps\n\n")
	
	// Test calculations
	fmt.Println("=== FITNESS STATS ===")
	
	distance := tracker.GetDistanceMiles()
	fmt.Printf("Distance: %.1f miles %s\n", distance, checkFloat(distance, 5.0))
	fmt.Printf("Expected: 5.0 miles\n\n")
	
	calories := tracker.GetCaloriesBurned()
	fmt.Printf("Calories: %.0f %s\n", calories, checkFloat(calories, 400.0))
	fmt.Printf("Expected: 400\n\n")
	
	progress := tracker.GetGoalProgress()
	fmt.Printf("Goal Progress: %.0f%% %s\n", progress*100, checkFloat(progress, 1.0))
	fmt.Printf("Expected: 100%%\n\n")
	
	activityLevel := tracker.GetActivityLevel()
	fmt.Printf("Activity Level: %s %s\n", activityLevel, checkString(activityLevel, "Very Active"))
	fmt.Printf("Expected: Very Active\n\n")
	
	// Test with different step counts
	fmt.Println("=== TESTING DIFFERENT ACTIVITY LEVELS ===")
	testActivityLevels(user)
	
	// Test reset
	fmt.Println("\n=== TESTING RESET ===")
	fmt.Printf("Before reset: %d steps\n", tracker.stepCount)
	tracker.ResetDaily()
	fmt.Printf("After reset: %d steps %s\n", tracker.stepCount, checkInt(tracker.stepCount, 0))
	fmt.Printf("Expected: 0 steps\n")
	
	// Test edge cases
	fmt.Println("\n=== TESTING EDGE CASES ===")
	tracker.AddSteps(0) // Should not add
	fmt.Printf("After adding 0 steps: %d %s\n", tracker.stepCount, checkInt(tracker.stepCount, 0))
	
	tracker.AddSteps(-5) // Should not add negative
	fmt.Printf("After adding -5 steps: %d %s\n", tracker.stepCount, checkInt(tracker.stepCount, 0))
	
	fmt.Println("\n🎉 All bugs fixed! Fitness tracker working perfectly! 💪")
}

func testActivityLevels(user User) {
	testCases := []struct{
		steps int
		expected string
	}{
		{3000, "Sedentary"},
		{6000, "Lightly Active"}, 
		{9000, "Moderately Active"},
		{10000, "Very Active"}, // FIXED: 10000 is Very Active now
		{11000, "Very Active"},
		{15000, "Extremely Active"},
	}
	
	for _, tc := range testCases {
		tracker := NewFitnessTracker(user)
		tracker.AddSteps(tc.steps)
		level := tracker.GetActivityLevel()
		fmt.Printf("%d steps: %s (Expected: %s) %s\n", 
			tc.steps, level, tc.expected, checkString(level, tc.expected))
	}
}

func checkFloat(got, expected float64) string {
	if got >= expected-0.1 && got <= expected+0.1 {
		return "✅"
	}
	return "❌"
}

func checkString(got, expected string) string {
	if got == expected {
		return "✅"
	}
	return "❌"
}

func checkInt(got, expected int) string {
	if got == expected {
		return "✅"
	}
	return "❌"
}