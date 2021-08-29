// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

package main

func findDisappearedNumbers(nums []int) []int {
    marked := make(map[int]bool)
    disNums := make([]int, 0)
    
    for _, val := range nums{
        if marked[val]{
            continue
        }
        marked[val] = true
    }
    
    for i:=1; i<len(nums)+1; i++ {
        if !marked[i]{
            disNums = append(disNums, i)
        }
    }
    
    return disNums
}