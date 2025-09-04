# Magic Potion Brewing - Solution 🧙‍♂️⚗️

## Approach

This is a character frequency counting problem, also known as the "Valid Anagram" or "Can Construct" problem. We need to check if the pantry contains at least the required quantities of each ingredient from the recipe.

## Algorithm

There are two main approaches:

### Approach 1: Hash Map (Character Counting)
1. **Count recipe ingredients:** Create a map of character frequencies for the recipe
2. **Count pantry ingredients:** Create a map of character frequencies for the pantry
3. **Compare:** Check if pantry has at least the required count for each recipe ingredient

### Approach 2: Consume from Pantry (More Memory Efficient)
1. **Count recipe ingredients:** Create a map of character frequencies for the recipe
2. **Iterate through pantry:** For each character in pantry, decrement its count in the recipe map
3. **Check remaining:** If any recipe ingredient count is still positive, return false

**Time Complexity:** O(n + m) where n = recipe length, m = pantry length
**Space Complexity:** O(k) where k = number of unique characters in recipe

## Key Edge Cases

- **Empty recipe:** Can always be brewed (return `true`)
- **Empty pantry with non-empty recipe:** Cannot be brewed (return `false`)
- **Duplicate ingredients:** Must count quantities correctly
- **Extra pantry ingredients:** Don't affect the result (can have more than needed)

## Step-by-step Example

For recipe `"love"` and pantry `"evolve"`:

1. Recipe needs: {l:1, o:1, v:1, e:1}
2. Process pantry:
   - 'e': {l:1, o:1, v:1, e:0} ✓
   - 'v': {l:1, o:1, v:0, e:0} ✓
   - 'o': {l:1, o:0, v:0, e:0} ✓
   - 'l': {l:0, o:0, v:0, e:0} ✓
   - 'v': already satisfied
   - 'e': already satisfied
3. All ingredients satisfied → return `true`