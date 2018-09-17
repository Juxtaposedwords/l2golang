package climb

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

So this is fibonacci and the hint is seen in the difference of steps, as is the
   exit case.

*/
var seen = map[int]int{0: 0, 1: 1, 2: 2, 3: 3}

func climbStairs(n int) int {
	_, ok := seen[n]
	if ok {
		return seen[n]
	}

	seen[n] = climbStairs(n-2) + climbStairs(n-1)
	return seen[n]
}
