# Coffee Shop Calculator ☕💰 [DEBUG ME!]

The local coffee shop's point-of-sale system is broken! Customers are getting incorrect totals and the manager needs your help to fix the calculator.

## The Problem

The coffee shop calculator is supposed to:
1. Calculate the total price for coffee orders
2. Apply a 10% discount for orders over $20
3. Add 8% sales tax to the final amount
4. Round to 2 decimal places

But something's wrong! Customers are complaining about incorrect charges.

## Test Cases That Should Work

```
Order 1: 3 Lattes ($4.50 each)
Expected: $14.58 (13.50 + 8% tax)
Getting: Wrong amount!

Order 2: 5 Cappuccinos ($5.00 each) 
Expected: $24.30 (25.00 - 10% discount + 8% tax)
Getting: Wrong amount!

Order 3: 1 Espresso ($2.00)
Expected: $2.16 (2.00 + 8% tax)
Getting: Wrong amount!
```

## Your Task

Find and fix ALL the bugs in the coffee shop calculator code. The main.go file contains several bugs that need to be identified and corrected.

## Hints
- Check the discount calculation logic
- Look at the tax calculation
- Verify the rounding function
- Check for edge cases

Debug the code and make customers happy again! ☕😊