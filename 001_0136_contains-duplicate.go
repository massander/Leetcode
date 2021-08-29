// https://leetcode.com/problems/contains-duplicate/

package main

func containsDuplicate(nums []int) bool {
	visited := make(map[int]int)
	for _, val := range nums {
		visited[val]++
		if visited[val] >= 2 {
			return true
		}
	}
	return false
}
