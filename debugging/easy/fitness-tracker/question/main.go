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

// BUG: This function has step counting bugs
func (ft *FitnessTracker) AddSteps(steps int) {
	if steps >= 0 { // BUG #1: Should be > 0, can't add 0 steps meaningfully
		ft.stepCount = steps // BUG #2: Should add to stepCount, not replace it
	}
}

// BUG: This function has distance calculation bug
func (ft *FitnessTracker) GetDistanceMiles() float64 {
	// Average person: 2000 steps = 1 mile
	return float64(ft.stepCount) * 2000.0 // BUG #3: Should divide by 2000, not multiply
}

// BUG: This function has calorie calculation bugs
func (ft *FitnessTracker) GetCaloriesBurned() float64 {
	// Formula: 4 calories per 100 steps for 150 lb person
	// Adjust for user weight: (userWeight / 150) * base calories
	
	baseCaloriesPer100Steps := 4.0
	stepsIn100s := float64(ft.stepCount) * 100 // BUG #4: Should divide by 100
	weightFactor := ft.user.WeightLbs + 150.0 // BUG #5: Should divide by 150, not add
	
	return stepsIn100s * baseCaloriesPer100Steps * weightFactor
}

// BUG: This function has goal progress calculation bug
func (ft *FitnessTracker) GetGoalProgress() float64 {
	if ft.user.DailyGoal == 0 {
		return 0.0
	}
	
	progress := float64(ft.stepCount) * float64(ft.user.DailyGoal) // BUG #6: Should divide, not multiply
	if progress >= 1.0 { // BUG #7: Should use > for exceeding 100%
		return 1.0 // Cap at 100%
	}
	return progress
}

// BUG: This function has activity level classification bugs  
func (ft *FitnessTracker) GetActivityLevel() string {
	steps := ft.stepCount
	
	if steps < 5000 {
		return "Sedentary"
	} else if steps < 7500 {
		return "Lightly Active"  
	} else if steps <= 10000 { // BUG #8: Should be < 10000 (not <=)
		return "Moderately Active"
	} else if steps < 12500 {
		return "Very Active"
	} else {
		return "Extremely Active"
	}
}

// BUG: This function has reset bug
func (ft *FitnessTracker) ResetDaily() {
	ft.stepCount = 1 // BUG #9: Should reset to 0, not 1
}

func main() {
	fmt.Println("🏃‍♂️ Fitness Tracker (BUGGY VERSION) 🏃‍♂️\n")
	
	// Create user and tracker
	user := User{
		Name:      "Alice",
		WeightLbs: 150.0,
		DailyGoal: 10000,
	}
	
	tracker := NewFitnessTracker(user)
	
	fmt.Printf("User: %s (%.0f lbs, Goal: %d steps)\\n\\n", user.Name, user.WeightLbs, user.DailyGoal)
	
	// Test adding steps
	fmt.Println("Adding 5000 steps...")
	tracker.AddSteps(5000)
	fmt.Printf("Steps: %d\\n", tracker.stepCount)
	fmt.Printf("Expected: 5000 steps\\n\\n")
	
	// Add more steps
	fmt.Println("Adding 3000 more steps...")
	tracker.AddSteps(3000)
	fmt.Printf("Steps: %d\\n", tracker.stepCount)
	fmt.Printf("Expected: 8000 steps\\n\\n")
	
	// Add final steps to reach 10,000
	fmt.Println("Adding 2000 more steps...")
	tracker.AddSteps(2000)
	fmt.Printf("Steps: %d\\n", tracker.stepCount)
	fmt.Printf("Expected: 10000 steps\\n\\n")
	
	// Test calculations
	fmt.Println("=== FITNESS STATS ===")
	
	distance := tracker.GetDistanceMiles()
	fmt.Printf("Distance: %.1f miles\\n", distance)
	fmt.Printf("Expected: 5.0 miles\\n")
	fmt.Printf("Status: %s\\n\\n", checkFloat(distance, 5.0))
	
	calories := tracker.GetCaloriesBurned()
	fmt.Printf("Calories: %.0f\\n", calories)
	fmt.Printf("Expected: 400\\n")
	fmt.Printf("Status: %s\\n\\n", checkFloat(calories, 400.0))
	
	progress := tracker.GetGoalProgress()
	fmt.Printf("Goal Progress: %.0f%%\\n", progress*100)
	fmt.Printf("Expected: 100%%\\n")
	fmt.Printf("Status: %s\\n\\n", checkFloat(progress, 1.0))
	
	activityLevel := tracker.GetActivityLevel()
	fmt.Printf("Activity Level: %s\\n", activityLevel)
	fmt.Printf("Expected: Very Active\\n")
	fmt.Printf("Status: %s\\n\\n", checkString(activityLevel, "Very Active"))
	
	// Test with different step counts
	fmt.Println("=== TESTING DIFFERENT ACTIVITY LEVELS ===")
	testActivityLevels(user)
	
	// Test reset
	fmt.Println("\\n=== TESTING RESET ===")
	fmt.Printf("Before reset: %d steps\\n", tracker.stepCount)
	tracker.ResetDaily()
	fmt.Printf("After reset: %d steps\\n", tracker.stepCount)
	fmt.Printf("Expected: 0 steps\\n")
	fmt.Printf("Status: %s\\n", checkInt(tracker.stepCount, 0))
}

func testActivityLevels(user User) {
	testCases := []struct{
		steps int
		expected string
	}{
		{3000, "Sedentary"},
		{6000, "Lightly Active"}, 
		{9000, "Moderately Active"},
		{11000, "Very Active"},
		{15000, "Extremely Active"},
	}
	
	for _, tc := range testCases {
		tracker := NewFitnessTracker(user)
		tracker.AddSteps(tc.steps)
		level := tracker.GetActivityLevel()
		fmt.Printf("%d steps: %s (Expected: %s) %s\\n", 
			tc.steps, level, tc.expected, checkString(level, tc.expected))
	}
}

func checkFloat(got, expected float64) string {
	if got >= expected-0.1 && got <= expected+0.1 {
		return "✅ PASS"
	}
	return "❌ FAIL"
}

func checkString(got, expected string) string {
	if got == expected {
		return "✅ PASS"
	}
	return "❌ FAIL"
}

func checkInt(got, expected int) string {
	if got == expected {
		return "✅ PASS"
	}
	return "❌ FAIL"
}