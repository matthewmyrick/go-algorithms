# Weather Reporter 🌤️📺 [DEBUG ME!]

The local weather station's reporting system is giving incorrect forecasts! Viewers are getting wrong temperature conversions, bad weather alerts, and confusing summaries.

## The Problem

The weather system should:
1. Convert between Fahrenheit and Celsius correctly
2. Determine appropriate weather conditions based on temperature
3. Generate accurate weather alerts for extreme conditions
4. Calculate weekly temperature averages
5. Provide helpful weather recommendations

But meteorologists are getting complaints about wrong weather data!

## Expected Behavior

```
Temperature Conversions:
- 32°F should be 0°C
- 0°C should be 32°F
- 100°F should be 37.78°C (rounded to 2 decimal places)

Weather Conditions:
- Below 32°F: "Freezing"
- 32-59°F: "Cold" 
- 60-79°F: "Pleasant"
- 80-89°F: "Warm"
- 90°F+: "Hot"

Weather Alerts:
- Below 20°F or above 100°F: Issue alert
- Otherwise: No alert needed
```

## Your Task

Find and fix ALL the bugs in the weather reporting system. The code has multiple issues with:
- Temperature conversion formulas
- Weather condition classification
- Alert generation logic
- Average calculation
- String formatting and comparisons

Help the meteorologists provide accurate weather reports! 🌡️⛅