// https://leetcode.com/problems/longest-palindromic-substring/

package main

func longestPalindrome(s string) string {
	var radius int
	var longestPalindrome string

	for i := 0; i < len(s); i++ {
		radius = 1
		for i-radius >= 0 && i+radius < len(s) && s[i-radius] == s[i+radius] {
			radius++
		}

		if len(s[i-radius+1:i+radius]) > len(longestPalindrome) {
			longestPalindrome = s[i-radius+1 : i+radius]
		}

		radius = 0
		for i-radius-1 >= 0 && i+radius < len(s) && s[i-radius-1] == s[i+radius] {
			radius++
		}

		if len(s[i-radius:i+radius]) > len(longestPalindrome) {
			longestPalindrome = s[i-radius : i+radius]
		}
	}

	return longestPalindrome
}
