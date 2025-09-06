# Movie Ticket System 🎬🎫 [DEBUG ME!]

The movie theater's ticket booking system is malfunctioning! Customers can't book seats properly and the system is showing incorrect availability information.

## The Problem

The movie ticket system should:
1. Allow customers to book available seats
2. Prevent double booking of the same seat
3. Show correct seat availability
4. Handle cancellations properly
5. Display the seating chart correctly

But there are several bugs causing issues!

## Expected Behavior

```
Theater: 5 rows x 5 seats each (A1-E5)
Available seats should show as "O"
Booked seats should show as "X" 

Booking "A1" should:
- Return true (successful booking)
- Mark A1 as unavailable
- Show A1 as "X" in seating chart

Booking same seat again should:
- Return false (already booked)
```

## Your Task

Find and fix ALL the bugs in the movie ticket system. The code has multiple issues with:
- Seat booking logic
- Availability checking  
- Seating chart display
- Cancellation functionality
- Index calculations

Debug the system and make movie-goers happy! 🍿😊