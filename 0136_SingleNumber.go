// https://leetcode.com/problems/single-number/

package main

func singleNumber(nums []int) int {
    marked := make(map[int]int)
    
    for _, val:= range nums{
        marked[val]++
    }
    
    for i, val:= range marked{
        if val == 1{
            return i
        }
    }

    return 0
}