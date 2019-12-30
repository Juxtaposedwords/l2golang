# Permutations

## Part 1
### Prompt (with clarification added)
```
--- Day 4: Secure Container ---
You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.

However, they do remember a few key facts about the password:

1. It is a six-digit number.
2. The value is within the range given in your puzzle input.
3. (There are at least) two adjacent digits (which) are the same (like 22 in 122345).
4. Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
* Other than the range rule, the following are true:
    * 111111 meets these criteria (double 11, never decreases).
    * 223450 does not meet these criteria (decreasing pair of digits 50).
    * 123789 does not meet these criteria (no double).

Question:
How many different passwords within the range given in your puzzle input meet these criteria?
```

### First Reading

* The problem is either:
    *  a straightforward computer science problem
        * **Brute Force(time:6^n space: n)**
            1.  For every number startign with the lowest possible combination and ending with the highest, issue a go routine which:
                a. Split the number into an array of ints
                b. Exit if there is no double or the array ever decreases
                c. Issue a counter
                    * Here we can use `sync/atomic`, which is hacky but works
        * **LookAhead(time: logN space: 1)**
            1. For every number between lower and upper. 
                a. check every digit:
                    * look at the next digit and compare it to the current:
                        i. **less** skip ahead at the highest level to the number where all numbers to the left are equal to the current 
                        ii. **greater** great
                        iii. **equal** Make a note you saw a repeat
                b. Increment if we saw a repeat and didn't have to skip


                    

    * an interesting combinatorics problem with a runtime of n
        * A solution without the range limits [is covered on math.stackechange](https://math.stackexchange.com/questions/336734/combinatorics-how-to-find-the-number-of-sets-of-numbers-in-increasing-order/336768#336768).  
        * I lack experience with combinatorics sadly, so it's speculation as to how tackle/estimate.


## Summary
