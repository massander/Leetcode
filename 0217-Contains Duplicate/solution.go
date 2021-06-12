package main

import "fmt"

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

func main(){
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
}