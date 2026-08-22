func twoSum(numbers []int, target int) []int {

	find := make(map[int]int)
	res := make([]int, 2)

	for i := range numbers {
		remainder := target - numbers[i]

		if j, ok := find[numbers[i]]; ok {
			res[0], res[1] = j, (i+1)
		}

		find[remainder] = (i + 1) 

	}

	return res
}
