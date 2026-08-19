func productExceptSelf(nums []int) []int {
	
	n := len(nums)

	prefix := make([]int, n)
	suffix := make([]int, n)
	prefix[0] = 1
	suffix[n-1] = 1

	// 1 3 6 24
	//[3,2,4,6]
	for i := 1; i < n; i++ {
		prefix[i] = nums[i-1] * prefix[i-1] 
	}
	// 48 24  6 1
	//[ 3, 2, 4,6]
	for i := n-2; i >= 0 ; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}

	for i := range nums {
		nums[i] = prefix[i] * suffix[i]
	}
	
	return nums
}
