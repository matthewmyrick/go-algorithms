# Coffee Shop Calculator - Solution ☕💰

## Bugs Found and Fixed

### Bug #1: Not accumulating subtotal
**Problem:** `subtotal = order.Price * float64(order.Quantity)`
**Fix:** `subtotal += order.Price * float64(order.Quantity)`
**Issue:** Was overwriting instead of adding to subtotal

### Bug #2: Incorrect discount condition  
**Problem:** `if subtotal >= 20`
**Fix:** `if subtotal > 20`
**Issue:** Discount should only apply to orders over $20, not including exactly $20

### Bug #3: Adding discount instead of subtracting
**Problem:** `discountedTotal = subtotal + discount`  
**Fix:** `discountedTotal = subtotal - discount`
**Issue:** Discount should reduce the total, not increase it

### Bug #4: Subtracting tax instead of adding
**Problem:** `finalTotal = discountedTotal - taxAmount`
**Fix:** `finalTotal = discountedTotal + taxAmount`  
**Issue:** Tax should be added to the total, not subtracted

### Bug #5: Incorrect decimal place rounding
**Problem:** `math.Round(finalTotal*10) / 10`
**Fix:** `math.Round(finalTotal*100) / 100`
**Issue:** For 2 decimal places, need to multiply/divide by 100, not 10

## Logic Flow (Corrected)

1. **Calculate subtotal:** Sum up all item prices × quantities
2. **Apply discount:** If subtotal > $20, subtract 10% discount  
3. **Add tax:** Add 8% sales tax to discounted total
4. **Round:** Round to 2 decimal places

## Common Debugging Strategies

- **Test edge cases:** Exactly $20 order, single items, multiple items
- **Print intermediate values:** Add debug prints to see subtotal, discount, tax
- **Check arithmetic:** Verify calculations manually
- **Read requirements carefully:** "Over $20" vs "at least $20"
- **Test incrementally:** Fix one bug at a time and retest