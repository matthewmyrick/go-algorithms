# Space Station Docking Sequence - Solution 🚀🛸

## Approach

This is a classic **Stack Sequence Validation** problem. We need to simulate the process of using a stack (holding orbit) to transform one sequence into another.

## Algorithm

We simulate the docking process:

1. **Initialize** an empty stack (holding orbit) and pointers for arrival and target sequences
2. **For each spacecraft in the arrival sequence:**
   - Push it onto the stack (enter holding orbit)
   - **While** the stack is not empty AND the top matches the next target spacecraft:
     - Pop from stack (spacecraft docks)
     - Move to next target spacecraft
3. **After processing all arrivals:** Check if stack is empty and all targets are achieved

**Key insight:** If at any point we can't match the next target spacecraft, either by pushing a new arrival or popping from the stack, the sequence is impossible.

**Time Complexity:** O(n) - each spacecraft is pushed and popped at most once
**Space Complexity:** O(n) - worst case, all spacecraft in holding orbit at once

## Step-by-step Example

For arrival `[1, 2, 3, 4, 5]` and target `[3, 2, 4, 1, 5]`:

1. **Process 1:** stack=[1], target_idx=0 (need 3)
2. **Process 2:** stack=[1,2], target_idx=0 (need 3)  
3. **Process 3:** stack=[1,2,3], target_idx=0 (need 3)
   - Pop 3: stack=[1,2], target_idx=1 (need 2) ✅
   - Pop 2: stack=[1], target_idx=2 (need 4) ✅
4. **Process 4:** stack=[1,4], target_idx=2 (need 4)
   - Pop 4: stack=[1], target_idx=3 (need 1) ✅
   - Pop 1: stack=[], target_idx=4 (need 5) ✅
5. **Process 5:** stack=[5], target_idx=4 (need 5)
   - Pop 5: stack=[], target_idx=5 (done) ✅

**Result:** Stack empty and all targets achieved → **Possible!**

## Key Learning Points

- **Stack simulation** is powerful for sequence validation problems
- **Greedy approach** works: always try to match the current target
- **Two pointers technique** efficiently tracks progress through sequences