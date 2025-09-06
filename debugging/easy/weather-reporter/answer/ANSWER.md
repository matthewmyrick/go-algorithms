# Weather Reporter - Solution 🌤️📺

## Bugs Found and Fixed

### Bug #1: Wrong Fahrenheit to Celsius conversion
**Problem:** `(fahrenheit + 32) * 5 / 9` (adding 32)
**Fix:** `(fahrenheit - 32) * 5 / 9` (subtracting 32)
**Issue:** Formula should subtract 32, not add

### Bug #2: Wrong decimal rounding
**Problem:** `math.Round(celsius*100) * 100` (multiplying by 100)
**Fix:** `math.Round(celsius*100) / 100` (dividing by 100) 
**Issue:** To round to 2 decimals: multiply, round, then divide by 100

### Bug #3: Wrong Celsius to Fahrenheit conversion
**Problem:** `celsius * 9 / 5 - 32` (subtracting 32)
**Fix:** `celsius * 9 / 5 + 32` (adding 32)
**Issue:** Formula should add 32, not subtract

### Bug #4: Wrong freezing point boundary
**Problem:** `if fahrenheit <= 32` (includes 32°F)
**Fix:** `if fahrenheit < 32` (below 32°F only)
**Issue:** 32°F is freezing point, but should be "Cold", not "Freezing"

### Bug #5: Range overlap issue 
**Problem:** Conditions had overlapping boundaries
**Fix:** Proper range boundaries: <32, 32-59, 60-79, etc.
**Issue:** Need clear non-overlapping temperature ranges

### Bug #6: Wrong logical operator in alert
**Problem:** `fahrenheit < 20 && fahrenheit > 100` (AND impossible)
**Fix:** `fahrenheit < 20 || fahrenheit > 100` (OR condition)
**Issue:** Temperature can't be both below 20 AND above 100

### Bug #7: Wrong alert boundary
**Problem:** `if fahrenheit <= 20` (includes 20°F)
**Fix:** `if fahrenheit < 20` (below 20°F only)
**Issue:** Alert should be for temperatures below 20°F

### Bug #8: Missing return in alert function
**Problem:** Function sometimes didn't return a value
**Fix:** Added proper return statement
**Issue:** All code paths must return a value

### Bug #9: Wrong array length check
**Problem:** `if len(ws.readings) >= 7` (removes at 7 readings)
**Fix:** `if len(ws.readings) > 7` (removes after 7 readings)
**Issue:** Should keep exactly 7 readings, not remove at 7

### Bug #10: Array index out of bounds
**Problem:** `for i := 0; i <= len(ws.readings)` (includes length)
**Fix:** `for i := 0; i < len(ws.readings)` (excludes length)
**Issue:** Array indices go from 0 to len-1, not len

### Bug #11: Wrong average calculation
**Problem:** `sum * float64(len(ws.readings))` (multiplying)
**Fix:** `sum / float64(len(ws.readings))` (dividing)
**Issue:** Average requires division, not multiplication

### Bug #12: Wrong alert display condition
**Problem:** `if strings.Contains(alert, "No weather alerts")` (shows when no alerts)
**Fix:** `if !strings.Contains(alert, "No weather alerts")` (shows when there are alerts)
**Issue:** Should show alerts when there ARE alerts, not when there aren't

### Bug #13: Wrong readings check
**Problem:** `if len(ws.readings) >= 0` (always true)
**Fix:** `if len(ws.readings) > 0` (has readings)
**Issue:** Should only show average when there are recorded temperatures

### Bug #14: Missing default case return
**Problem:** Switch case had no default return
**Fix:** Added return statement in default case
**Issue:** All code paths must return a value

## Key Learning Points

- **Formula accuracy:** Temperature conversion formulas must be exact
- **Boundary conditions:** Careful with inclusive vs exclusive ranges  
- **Logical operators:** AND vs OR in conditional statements
- **Array bounds:** Indices from 0 to len-1, not len
- **Mathematical operations:** Division vs multiplication in calculations
- **Return statements:** All code paths must return appropriate values