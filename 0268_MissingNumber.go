// https://leetcode.com/problems/missing-number/

package main

func missingNumber(nums []int) int {
    marked := make(map[int]bool)
    for _, val := range nums{
        marked[val] = true
    }
    
    for i:=0; i<len(nums)+1; i++ {
        if !marked[i]{
            return i
        }
    }
    
    return 0
}