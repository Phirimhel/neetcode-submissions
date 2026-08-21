func longestConsecutive(nums []int) int {

	n := len(nums)
	if n == 1 {
		return 1 
	}
	if n == 0 {
		return 0
	}
	sort.Ints(nums)

	counter := 1
	sequence := 1 

	//[-1 -1 0 1 3 4 5 6 7 8 9]
	for i := 0; i < (n -1); i++ {	

		if nums[i] == nums[i+1] {
			continue
		}

		if nums[i]+1 == nums[i+1] {
			counter++ 
		} else {
			counter = 1
		}

		if counter > sequence {
			sequence = counter 
		}

		

		
	}

	return sequence
}
