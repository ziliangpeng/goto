func twoSum(nums []int, target int) []int {
    for i, val1 := range nums {
        for j, val2 := range nums[i+1:] {
            if val1 + val2 == target {
                return []int{i, i+1+j}
            }
        }
    }
	return []int{0, 0}
}
