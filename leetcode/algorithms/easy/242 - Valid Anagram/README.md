# Intuition
A naive approach is to sort, but that's not how we thinkg about this. We know an optimal solution is one which will have to read through the first word fully, but may be able to skip the second word. 

We see that when we solve for the word 'foooooo' and 'fog': when we get to the second word. 


# Approach
Put word one's rune count in a map. Decrement from that map as we read through the second one and return early if find an unknown match.

This gives us the good fortune of handling of rune count mismatch (either 2 vs 3 or 1 vs 0).

We'll cheat and add the old classic if lengths are not equal up front to optimize.

# Complexity
- Time complexity:
  
Big O: N of S (how many runes in S) + N of T (how many runes in T). Let's call it N between friends O(n).

- Space complexity:
Big O: N of S (how many runes in S) + N of T (how many runes in T). Let's call it N between friends O(n).

# Code
```
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false // acounts for beating 20% of time!
    }
	sRunes := map[rune]int{}
	for _, r := range s {
		sRunes[r]++
	}
	for _, r := range t {
        // This part really saves us heartache as now we're not play kooky 'is it -1?' games.
		if _, ok := sRunes[r]; !ok {
			return false
		}
		sRunes[r]--
		if sRunes[r] == 0 {
			delete(sRunes, r)
		}
	}
	return len(sRunes) == 0
}

```

## Thoughts
Overall I'm sure there are some cool and faster solutions, but i'm happy with the peformance and the straight-forwardness of this solution.