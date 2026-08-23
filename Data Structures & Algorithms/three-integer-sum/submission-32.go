func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	output := make([][]int, 0, len(nums))
	length := len(nums) - 1 
	
	
	// -4 -1 -1 0 1 2 
	//        i 
	//          k   j 


	//  -1 -1 0 1 
	//      i 
	//        k j


	for i := range nums {

		if i > 0 && nums[i] == nums[i-1] {
				continue
		}

		k := i + 1 
		j := length
		 
		for k < j {
			sum := nums[i] + nums[k] + nums[j]

			if sum == 0 {
				output = append(output, []int{nums[i], nums[k], nums[j]})
				k++
				j--

    			for k < j && nums[k] == nums[k-1] {
        			k++
    			}
    			for k < j && nums[j] == nums[j+1] {
        			j--
    			}
			}

			if sum < 0 {
				k++
			}

			if sum > 0 {
				j--
			}
			
		}

	}
	return output
}



