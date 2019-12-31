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

### Notes

* The problem is either:
    *  a straightforward computer science problem
        * **Brute Force(time:m (where m is the numbers between lower and upper bound) space: constant)**
            1.  For every number starting with the lowest possible combination and ending with the highest, issue a go routine which:
                a. Split the number into an array of ints
                b. Exit if there is no double or the array ever decreases
                c. Issue a counter
                    * Here we can use `sync/atomic`, which is hacky but works
        * **LookAhead(time: n^2 (where n is the number of digits) space: constant)**
            1. For every number between lower and upper. 
                a. check every digit:
                    * look at the next digit and compare it to the current:
                        i. **less** skip ahead at the highest level to the number where all numbers to the left are equal to the current 
                        ii. **greater** great
                        iii. **equal** Make a note you saw a repeat
                b. Increment if we saw a repeat and didn't have to skip
    * an interesting combinatorics problem with a runtime of sub logN (much like Binet's formula for Fibonacci)
        * A solution without the range limits [is covered on math.stackechange](https://math.stackexchange.com/questions/336734/combinatorics-how-to-find-the-number-of-sets-of-numbers-in-increasing-order/336768#336768).  
        * I lack experience with combinatorics sadly, so it's speculation as to how tackle/estimate.

## Part 2

### Prompt 
```
--- Part Two ---
An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of matching digits.

Given this additional criterion, but still ignoring the range rule, the following are now true:

112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
How many different passwords within the range given in your puzzle input meet all of the criteria?
```

### Notes
* The author says that anything repeat that is in a group larger than two does not count as our double. By this reasoning one is fine, two is not, as 4 is in a group of 3.
* Since we check for increasing earlier in our vetting process, we can:
    1. record each time we see a digit. 
    2. check if we ever a digit which occured twice. 