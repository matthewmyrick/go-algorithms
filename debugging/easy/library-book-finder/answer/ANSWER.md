# Library Book Finder - Solution 📚🔍

## Bugs Found and Fixed

### Bug #1: Inconsistent case conversion in search
**Problem:** `query = strings.ToUpper(query)`
**Fix:** `query = strings.ToLower(query)`
**Issue:** Query converted to upper but compared against lowercase book fields

### Bug #2 & #3: Wrong search logic (AND instead of OR)
**Problem:** `strings.Contains(title, query) && strings.Contains(author, query) && strings.Contains(ISBN, query)`
**Fix:** `strings.Contains(title, query) || strings.Contains(author, query) || strings.Contains(ISBN, query)`
**Issue:** Should find books matching ANY field, not ALL fields

### Bug #4: Wrong availability check in checkout
**Problem:** `if l.books[i].IsAvailable == false` (checking if unavailable)
**Fix:** `if !l.books[i].IsAvailable` (checking if unavailable, but logic was backwards)
**Issue:** Should only checkout available books

### Bug #5: Wrong availability status after checkout
**Problem:** `l.books[i].IsAvailable = true` (marking as available)
**Fix:** `l.books[i].IsAvailable = false` (marking as unavailable)
**Issue:** Checked out books should be unavailable

### Bug #6: Wrong loan period
**Problem:** `time.Now().AddDate(0, 0, 7)` (7 days)
**Fix:** `time.Now().AddDate(0, 0, 14)` (14 days)
**Issue:** Library loan period should be 14 days, not 7

### Bug #7: Wrong availability check in return
**Problem:** `if l.books[i].IsAvailable` (checking if available)
**Fix:** `if l.books[i].IsAvailable` (correct check)
**Issue:** Should only return books that are checked out (unavailable)

### Bug #8: Wrong availability status after return
**Problem:** `l.books[i].IsAvailable = false` (marking as unavailable)
**Fix:** `l.books[i].IsAvailable = true` (marking as available)
**Issue:** Returned books should be available

### Bug #9: Wrong patron field clearing
**Problem:** `l.books[i].CheckedOutBy = "RETURNED"`
**Fix:** `l.books[i].CheckedOutBy = ""`
**Issue:** Should clear patron name, not set to "RETURNED"

### Bug #10: Wrong due date clearing
**Problem:** `l.books[i].DueDate = time.Now()`
**Fix:** `l.books[i].DueDate = time.Time{}`
**Issue:** Should clear due date to zero time, not current time

### Bug #11: Wrong overdue days calculation
**Problem:** `time.Now().Sub(book.DueDate).Hours() * 24`
**Fix:** `time.Now().Sub(book.DueDate).Hours() / 24`
**Issue:** Hours to days requires division, not multiplication

### Bug #12: Wrong empty results check
**Problem:** `if len(books) > 0` then show "No books found"
**Fix:** `if len(books) == 0` then show "No books found"
**Issue:** Logic was backwards - should check for empty results

### Bug #13 & #14: Backwards status display logic
**Problem:** Available shows as "Checked out", details shown for available books
**Fix:** Available shows as "Available", details shown for checked out books
**Issue:** Display logic was completely backwards

## Key Learning Points

- **Logical operators:** AND vs OR in search conditions
- **Boolean state management:** Consistent meaning of true/false across functions
- **String comparison:** Consistent case handling in searches
- **Date arithmetic:** Proper conversion between time units
- **State transitions:** Proper field updates during status changes