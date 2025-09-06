package main

import (
	"fmt"
	"math"
	"strings"
)

type WeatherStation struct {
	location string
	readings []float64 // Temperature readings in Fahrenheit
}

func NewWeatherStation(location string) *WeatherStation {
	return &WeatherStation{
		location: location,
		readings: make([]float64, 0),
	}
}

// FIXED: Temperature conversion corrected
func (ws *WeatherStation) FahrenheitToCelsius(fahrenheit float64) float64 {
	// Formula: C = (F - 32) * 5/9
	celsius := (fahrenheit - 32) * 5 / 9 // FIXED: Subtract 32, not add
	return math.Round(celsius*100) / 100  // FIXED: Divide by 100 for 2 decimal places
}

// FIXED: Temperature conversion corrected
func (ws *WeatherStation) CelsiusToFahrenheit(celsius float64) float64 {
	// Formula: F = C * 9/5 + 32
	fahrenheit := celsius * 9 / 5 + 32 // FIXED: Add 32, not subtract
	return math.Round(fahrenheit*100) / 100
}

// FIXED: Weather condition boundaries corrected
func (ws *WeatherStation) GetWeatherCondition(fahrenheit float64) string {
	if fahrenheit < 32 { // FIXED: Use < for below 32
		return "Freezing"
	} else if fahrenheit >= 32 && fahrenheit <= 59 { // FIXED: Clear range 32-59
		return "Cold"
	} else if fahrenheit >= 60 && fahrenheit <= 79 {
		return "Pleasant"
	} else if fahrenheit >= 80 && fahrenheit <= 89 {
		return "Warm"
	} else {
		return "Hot"
	}
}

// FIXED: Weather alert logic corrected
func (ws *WeatherStation) GenerateWeatherAlert(fahrenheit float64) string {
	if fahrenheit < 20 || fahrenheit > 100 { // FIXED: Use OR (||), not AND (&&)
		if fahrenheit < 20 { // FIXED: Use < 20, not <= 20
			return "EXTREME COLD WARNING: Temperature below 20°F!"
		} else {
			return "EXTREME HEAT WARNING: Temperature above 100°F!"
		}
	}
	return "No weather alerts." // FIXED: Proper return for normal conditions
}

// FIXED: Temperature recording corrected
func (ws *WeatherStation) RecordTemperature(fahrenheit float64) {
	ws.readings = append(ws.readings, fahrenheit)
	
	// Remove old readings to keep only last 7 days
	if len(ws.readings) > 7 { // FIXED: Use > 7, not >= 7
		ws.readings = ws.readings[1:] // Remove first element
	}
}

// FIXED: Average calculation corrected
func (ws *WeatherStation) GetWeeklyAverage() float64 {
	if len(ws.readings) == 0 {
		return 0.0
	}
	
	var sum float64
	for i := 0; i < len(ws.readings); i++ { // FIXED: Use < len, not <= len
		sum += ws.readings[i]
	}
	
	return sum / float64(len(ws.readings)) // FIXED: Divide, not multiply
}

// FIXED: Report generation corrected
func (ws *WeatherStation) GenerateReport(fahrenheit float64) string {
	celsius := ws.FahrenheitToCelsius(fahrenheit)
	condition := ws.GetWeatherCondition(fahrenheit)
	alert := ws.GenerateWeatherAlert(fahrenheit)
	
	report := fmt.Sprintf("Weather Report for %s\n", ws.location)
	report += fmt.Sprintf("Temperature: %.1f°F (%.1f°C)\n", fahrenheit, celsius)
	report += fmt.Sprintf("Condition: %s\n", condition)
	
	if !strings.Contains(alert, "No weather alerts") { // FIXED: Show when there ARE alerts
		report += fmt.Sprintf("Alert: %s\n", alert)
	}
	
	if len(ws.readings) > 0 { // FIXED: Check > 0, not >= 0
		avgTemp := ws.GetWeeklyAverage()
		report += fmt.Sprintf("Weekly Average: %.1f°F\n", avgTemp)
	}
	
	return report
}

// FIXED: Clothing recommendation corrected
func (ws *WeatherStation) GetClothingRecommendation(fahrenheit float64) string {
	condition := ws.GetWeatherCondition(fahrenheit)
	
	switch condition {
	case "Freezing":
		return "Wear heavy winter coat, gloves, and hat"
	case "Cold":
		return "Wear jacket or sweater"
	case "Pleasant":
		return "Light clothing is perfect"
	case "Warm":
		return "Wear light, breathable clothing"
	case "Hot":
		return "Wear minimal, cooling clothing and stay hydrated"
	default: // FIXED: Added proper default return
		return "Weather condition unknown - dress appropriately"
	}
}

func main() {
	fmt.Println("🌤️ Weather Reporter (FIXED VERSION) 🌤️\n")
	
	station := NewWeatherStation("Downtown Weather Station")
	
	// Test temperature conversions
	fmt.Println("=== TEMPERATURE CONVERSION TESTS ===")
	
	// Test F to C
	celsius32 := station.FahrenheitToCelsius(32)
	fmt.Printf("32°F = %.2f°C (Expected: 0.00°C) %s\n", celsius32, checkFloat(celsius32, 0.0))
	
	celsius100 := station.FahrenheitToCelsius(100)
	fmt.Printf("100°F = %.2f°C (Expected: 37.78°C) %s\n", celsius100, checkFloat(celsius100, 37.78))
	
	// Test C to F
	fahrenheit0 := station.CelsiusToFahrenheit(0)
	fmt.Printf("0°C = %.2f°F (Expected: 32.00°F) %s\n", fahrenheit0, checkFloat(fahrenheit0, 32.0))
	
	fahrenheit20 := station.CelsiusToFahrenheit(20)
	fmt.Printf("20°C = %.2f°F (Expected: 68.00°F) %s\n\n", fahrenheit20, checkFloat(fahrenheit20, 68.0))
	
	// Test weather conditions
	fmt.Println("=== WEATHER CONDITION TESTS ===")
	testTemps := []float64{25, 45, 70, 85, 95}
	expectedConditions := []string{"Freezing", "Cold", "Pleasant", "Warm", "Hot"}
	
	for i, temp := range testTemps {
		condition := station.GetWeatherCondition(temp)
		fmt.Printf("%.0f°F: %s (Expected: %s) %s\n", 
			temp, condition, expectedConditions[i], checkString(condition, expectedConditions[i]))
	}
	
	// Test weather alerts
	fmt.Println("\n=== WEATHER ALERT TESTS ===")
	
	alert15 := station.GenerateWeatherAlert(15)
	fmt.Printf("15°F alert: %s %s\n", alert15, checkStringContains(alert15, "EXTREME COLD"))
	
	alert105 := station.GenerateWeatherAlert(105)
	fmt.Printf("105°F alert: %s %s\n", alert105, checkStringContains(alert105, "EXTREME HEAT"))
	
	alert70 := station.GenerateWeatherAlert(70)
	fmt.Printf("70°F alert: %s %s\n\n", alert70, checkStringContains(alert70, "No weather alerts"))
	
	// Test temperature recording and averaging
	fmt.Println("=== TEMPERATURE RECORDING TESTS ===")
	
	// Record some temperatures
	temps := []float64{65, 68, 72, 70, 75, 73, 71}
	for _, temp := range temps {
		station.RecordTemperature(temp)
	}
	
	average := station.GetWeeklyAverage()
	expectedAvg := 70.57 // (65+68+72+70+75+73+71)/7
	fmt.Printf("Weekly average: %.2f°F (Expected: %.2f°F) %s\n\n", average, expectedAvg, checkFloat(average, expectedAvg))
	
	// Test array length management
	fmt.Println("Recorded temperatures:", station.readings)
	fmt.Printf("Number of readings: %d (Expected: 7) %s\n\n", len(station.readings), checkInt(len(station.readings), 7))
	
	// Test full weather report
	fmt.Println("=== WEATHER REPORT TEST ===")
	report := station.GenerateReport(85)
	fmt.Println(report)
	
	// Test clothing recommendations
	fmt.Println("=== CLOTHING RECOMMENDATION TESTS ===")
	recommendation := station.GetClothingRecommendation(85)
	fmt.Printf("85°F clothing: %s %s\n", recommendation, checkStringContains(recommendation, "light"))
	
	// Test edge case: exactly at boundary temperatures
	fmt.Println("\n=== BOUNDARY CONDITION TESTS ===")
	fmt.Printf("32°F condition: %s (Expected: Cold) %s\n", 
		station.GetWeatherCondition(32), checkString(station.GetWeatherCondition(32), "Cold"))
	fmt.Printf("20°F alert: %s (Expected: No alerts) %s\n", 
		station.GenerateWeatherAlert(20), checkStringContains(station.GenerateWeatherAlert(20), "No weather"))
	
	fmt.Println("\n🎉 All bugs fixed! Weather system reporting accurately! 🌡️")
}

func checkFloat(got, expected float64) string {
	if math.Abs(got-expected) < 0.1 {
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

func checkStringContains(got, expected string) string {
	if strings.Contains(strings.ToLower(got), strings.ToLower(expected)) {
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