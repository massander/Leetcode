func twoSum(nums []int, target int) []int {
	// output := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if (i != j) && (nums[i]+nums[j] == target) {
				// output = append(output, i, j)
				return []int{i, j}
			}
		}
	}
	// return output
	return []int{}
}