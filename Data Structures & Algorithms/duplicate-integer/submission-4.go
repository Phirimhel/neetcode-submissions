func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{}, len(nums))

	for i := range nums {
		if _, ok := seen[nums[i]]; ok {
			return true
		}
		seen[nums[i]] = struct{}{}
	} 

	return false 
}
