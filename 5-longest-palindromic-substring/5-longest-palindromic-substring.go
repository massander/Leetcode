func longestPalindrome(s string) string {
	d1, d2 := make([]int, len(s)), make([]int, len(s))
	var longestPalindrome string = string(s[0])

	for i := 0; i < len(s); i++ {
		d1[i] = 1
		for i-d1[i] >= 0 && i+d1[i] < len(s) && s[i-d1[i]] == s[i+d1[i]] {

			d1[i]++
		}

		if len(s[i-d1[i]+1:i+d1[i]]) > len(longestPalindrome) {
			longestPalindrome = s[i-d1[i]+1 : i+d1[i]]
		}

		d2[i] = 0
		for i-d2[i]-1 >= 0 && i+d2[i] < len(s) && s[i-d2[i]-1] == s[i+d2[i]] {
			d2[i]++
		}

		if len(s[i-d2[i]:i+d2[i]]) > len(longestPalindrome) {
			longestPalindrome = s[i-d2[i] : i+d2[i]]
		}
	}

	return longestPalindrome
}
