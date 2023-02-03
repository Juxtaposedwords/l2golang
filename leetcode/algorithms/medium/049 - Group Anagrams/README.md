# Group Anagarams
## Problem statement

Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]
 

## Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.


## First thoughts 

Well I know an anagram can be identified and sorted if stored in an order representation of character->count. 

So my first approach would be to:
1. For each word:
   * Get the count of characters and store as a string ordered lexographically by letter, example: "eat"->"a1e1t1"
      * use this as dictionary key and append to the array the result.
2. Iterate over the dictionary, and return the different results.

Looking at the question just generally, it feels like a 'heaps' question though i'm not able to grok how i'd use them here. 

Run time: I'm thinking an O(n) run time as we build a dictionary and iterate rather than trying to sort. 
