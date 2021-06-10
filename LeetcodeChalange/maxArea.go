package main

import (
	"sort"
)

// 76ms/8.8MB
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

// 80ms/8.9MB
// func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
// 	sort.Ints(horizontalCuts)
// 	sort.Ints(verticalCuts)

// 	return findMaxDiffCut(h, horizontalCuts) * findMaxDiffCut(w, verticalCuts) % 1000000007
// }

// func findMaxDiffCut(size int, cuts []int) int {
// 	maxDiff := cuts[0]
// 	for i := 1; i < len(cuts); i++ {
// 		if cuts[i]-cuts[i-1] > maxDiff {
// 			maxDiff = cuts[i] - cuts[i-1]
// 		}
// 	}

// 	if size-cuts[len(cuts)-1] > maxDiff {
// 		return size - cuts[len(cuts)-1]
// 	}

// 	return maxDiff
// }
