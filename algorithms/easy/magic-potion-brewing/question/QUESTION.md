# Magic Potion Brewing 🧙‍♂️⚗️

You're an apprentice wizard learning to brew magic potions! Your master has given you a recipe, but the ingredients are all mixed up. You need to find if you have all the required ingredients in your pantry, regardless of their order.

## The Problem

Given two strings:
1. **Recipe** - the ingredients needed for the potion (each character represents an ingredient)
2. **Pantry** - the ingredients you have available

Determine if you can brew the potion by checking if the pantry contains all ingredients from the recipe with at least the required quantities.

## Example

**Recipe:** `"love"`
**Pantry:** `"evolve"`

**Analysis:**
- Need: l(1), o(1), v(1), e(1)
- Have: e(2), v(1), o(1), l(1)
- Result: ✅ **Yes!** You can brew the potion!

**Recipe:** `"magic"`
**Pantry:** `"gem"`

**Analysis:**
- Need: m(1), a(1), g(1), i(1), c(1)
- Have: g(1), e(1), m(1)
- Result: ❌ **No!** You're missing 'a', 'i', and 'c'

## Your Task

Write a function that takes a recipe string and a pantry string, and returns `true` if you can brew the potion, `false` otherwise.

**Function signature:**
```go
func CanBrewPotion(recipe, pantry string) bool
```

## Test Cases

1. `("love", "evolve")` → `true`
2. `("magic", "gem")` → `false`
3. `("aab", "baa")` → `true`
4. `("aab", "ab")` → `false` (need 2 'a's but only have 1)
5. `("", "anything")` → `true` (empty recipe needs no ingredients)
6. `("something", "")` → `false` (can't brew with empty pantry)

May your potions be powerful and your ingredients plentiful! 🌟✨