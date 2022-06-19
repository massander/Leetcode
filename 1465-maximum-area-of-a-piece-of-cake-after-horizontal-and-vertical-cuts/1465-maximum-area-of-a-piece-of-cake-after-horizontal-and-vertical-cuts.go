func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxHorizontalDiff := max(horizontalCuts[0], h-horizontalCuts[len(horizontalCuts)-1])
	maxVerticalDiff := max(verticalCuts[0], w-verticalCuts[len(verticalCuts)-1])

	for i := 1; i < len(horizontalCuts); i++ {
		maxHorizontalDiff = max(maxHorizontalDiff, horizontalCuts[i]-horizontalCuts[i-1])
	}

	for i := 1; i < len(verticalCuts); i++ {
		maxVerticalDiff = max(maxVerticalDiff, verticalCuts[i]-verticalCuts[i-1])
	}

	return (maxHorizontalDiff * maxVerticalDiff) % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
