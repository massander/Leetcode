// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

func lengthOfLongestSubstring(s string) int {
	var maxLen int = 0
	var lastIndex map[string]int = make(map[string]int)
	var startIndex int = 0
	for i := 0; i < len(s); i++ {
		if _, ok := lastIndex[string(s[i])]; ok {
			startIndex = Max(startIndex, lastIndex[string(s[i])]+1)
		}

		maxLen = Max(maxLen, i-startIndex+1)

		lastIndex[string(s[i])] = i
	}
	return maxLen
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
