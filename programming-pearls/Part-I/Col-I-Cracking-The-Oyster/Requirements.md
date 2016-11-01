Specifications:
- The file contains 27,000 records
- Each record is a 16 bit word
- Each record is in the range 1..27,000 
- No integer can appear more than once

Actual Problem:
The problem is one of political redistricting. The numbers to be sorted were indices of precincts that make up a political district.  Each district had a unique integer identifier from 1 to 27,000.
The desired output was a list of precinct numbers in a district sorted in numeric order.  

A user will call this once an hour and was a blocker for work. So the sort can't take more than minutes.

Defining the problem:

Input: a file with 27,000 unique integers in the range 1..27,000. Duplicates are an error.

Output: a sorted list increasing order of the input integers.

Constraints: at most there can be one thousand 16-bit words available in main memory (disk storage and buffer are infinite)

Problem:
How do you present 27,000 distinct integers in about 16,00 available bits?
