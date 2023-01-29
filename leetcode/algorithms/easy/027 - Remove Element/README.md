# Remove Element

* Given array `nums` and scalar `val`, remove all occurences of `val` from `nums` in place, without distorting the ordering.
   * The first K elements of nums should be the values.
   * Return K after placing the final result of the first k slots of nums.
   *  Must be performed in place

## Function signature
```
func removeElement(nums []int, val int) int {
    
}
```

## Provided test cases

```
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
```

## Constraings
* 0 <= nums.length <= 100
* 0 <= nums[i] <= 50
* 0 <= val <= 100

## Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]


# Thoughts:

* We don't _have_ to rearrange that slice input. They can't check.
* Looks like a common enough pattern of two indexes (one for last seen and one for counting). 
   * So we'll have an i which tracks our progress through the array while j tracks the last point in which we should assign something
       * We assign if i!=j and stop when i hits the max