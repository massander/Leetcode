// https://leetcode.com/problems/climbing-stairs/

package main

func climbStairs(n int) int {
	if n < 4 {
		return n
	}

	pre1 := 1
	pre2 := 2

	for i := 3; i <= n; i++ {
		pre1, pre2 = pre2, pre1+pre2
	}

	return pre2
}
