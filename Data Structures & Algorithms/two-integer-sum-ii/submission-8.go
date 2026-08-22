func twoSum(numbers []int, target int) []int {
	output := make([]int, 2)

	n := len(numbers)

	left := 0 
	right := n-1 
	for left <= right {

		sum := numbers[left] + numbers[right]

		if sum == target {
			output[0], output[1] = left + 1, right + 1
			break
		}

		if sum > target { 
			right--
			continue
		}

		if sum < target { 
			left++
			continue
		}

	}

	return output
}
