func threeSum(nums []int) [][]int {
	output := make([][]int, 0, len(nums))

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i + 1 && nums[j-1] == nums[j] {
				continue
			}

			for k := j + 1; k < len(nums); k++ {
				if k > j + 1 && nums[k-1] == nums[k] {
					continue
				}
				
				if nums[i] + nums[j] + nums[k] == 0 {
					output = append(output, []int{nums[i], nums[j],nums[k]})
				}
			}
		} 

	}

	return output
}
