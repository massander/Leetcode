type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	return NumArray{sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0{
		return this.nums[right] - this.nums[left-1]
	}

	return this.nums[right]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */