# Movie Ticket System - Solution 🎬🎫

## Bugs Found and Fixed

### Bug #1: Invalid seat validation logic
**Problem:** `if valid {` (wrong condition)
**Fix:** `if !valid {`  
**Issue:** Should return false when seat ID is invalid, not valid

### Bug #2: Not a bug, but could be clearer
**Current:** `if t.seats[row][col] == true`
**Better:** `if t.seats[row][col]` (more concise)
**Issue:** Redundant comparison, but functionally correct

### Bug #3: Wrong seat status assignment
**Problem:** `t.seats[row][col] = false` (marking as available)
**Fix:** `t.seats[row][col] = true` (marking as booked)
**Issue:** Booked seats should be true, not false

### Bug #4: Incorrect range validation logic  
**Problem:** `if row < 0 && row >= t.rows` (AND condition impossible)
**Fix:** `if row < 0 || row >= t.rows` (OR condition)
**Issue:** AND makes condition impossible to satisfy

### Bug #5: Wrong column index conversion
**Problem:** `col = col + 1` (making 1-based into 2-based)
**Fix:** `col = col - 1` (converting 1-based to 0-based)
**Issue:** Seat numbers are 1-based, but array indices are 0-based

### Bug #6: Wrong availability check
**Problem:** `return t.seats[row][col] == true` (returns booked status)
**Fix:** `return t.seats[row][col] == false` (returns available status)  
**Issue:** Available should mean not booked (false)

### Bug #7: Wrong cancellation status
**Problem:** `t.seats[row][col] = true` (marking as booked)
**Fix:** `t.seats[row][col] = false` (marking as available)
**Issue:** Cancelling should make seat available, not booked

### Bug #8 & #9: Swapped display symbols
**Problem:** `true` shows "O", `false` shows "X"
**Fix:** `true` shows "X" (booked), `false` shows "O" (available)
**Issue:** Display symbols were backwards

## Key Learning Points

- **Boolean logic:** AND vs OR conditions, especially in validation
- **Index conversion:** 1-based user input to 0-based array indices  
- **State representation:** Consistent meaning of true/false throughout system
- **Display logic:** Symbols should match internal representation
- **Edge case testing:** Invalid seats, double booking, cancelling unbooked seats