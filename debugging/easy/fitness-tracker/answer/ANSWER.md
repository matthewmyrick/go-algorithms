# Fitness Tracker - Solution 🏃‍♂️📱

## Bugs Found and Fixed

### Bug #1: Allowing zero steps
**Problem:** `if steps >= 0` (allows adding 0 steps)
**Fix:** `if steps > 0` 
**Issue:** Adding 0 steps shouldn't change step count

### Bug #2: Overwriting instead of accumulating steps
**Problem:** `ft.stepCount = steps` (overwrites total)
**Fix:** `ft.stepCount += steps` (adds to total)
**Issue:** Steps should accumulate throughout the day

### Bug #3: Wrong distance calculation
**Problem:** `float64(ft.stepCount) * 2000.0` (multiplying)
**Fix:** `float64(ft.stepCount) / 2000.0` (dividing)
**Issue:** 2000 steps = 1 mile, so divide steps by 2000

### Bug #4: Wrong steps-to-hundreds conversion
**Problem:** `float64(ft.stepCount) * 100` (multiplying)
**Fix:** `float64(ft.stepCount) / 100` (dividing) 
**Issue:** Need number of 100-step units, so divide by 100

### Bug #5: Wrong weight factor calculation
**Problem:** `ft.user.WeightLbs + 150.0` (adding)
**Fix:** `ft.user.WeightLbs / 150.0` (dividing)
**Issue:** Weight factor is a ratio: user_weight / base_weight

### Bug #6: Wrong goal progress calculation  
**Problem:** `float64(ft.stepCount) * float64(ft.user.DailyGoal)` (multiplying)
**Fix:** `float64(ft.stepCount) / float64(ft.user.DailyGoal)` (dividing)
**Issue:** Progress = steps_taken / goal_steps

### Bug #7: Wrong progress cap condition
**Problem:** `if progress >= 1.0` then cap at 100%
**Fix:** `if progress > 1.0` then cap at 100%
**Issue:** Should allow exactly 100% (1.0) without capping

### Bug #8: Wrong activity level boundary
**Problem:** `steps <= 10000` for "Moderately Active"
**Fix:** `steps < 10000` for "Moderately Active"
**Issue:** Exactly 10000 steps should be "Very Active", not "Moderately Active"

### Bug #9: Wrong reset value
**Problem:** `ft.stepCount = 1` (resets to 1)
**Fix:** `ft.stepCount = 0` (resets to 0)
**Issue:** Daily reset should start from 0 steps

## Key Learning Points

- **Accumulation vs Assignment:** Use `+=` for adding to totals
- **Unit Conversions:** Be careful about multiply vs divide in conversions
- **Ratio Calculations:** Weight factors and progress are ratios (division)
- **Boundary Conditions:** Inclusive vs exclusive ranges in classifications
- **Input Validation:** Consider edge cases like zero or negative inputs