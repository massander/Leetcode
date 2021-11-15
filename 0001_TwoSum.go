// https://leetcode.com/problems/two-sum/

package main

// // Bruteforce
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// return []int{}
// }

// Hash table
func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if _, ok := prevMap[diff]; ok {
			return []int{prevMap[diff], i}
		}
		prevMap[nums[i]] = i
	}
	return []int{}
}
