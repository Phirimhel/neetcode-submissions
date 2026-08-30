func dailyTemperatures(temperatures []int) []int {

	output := make([]int, len(temperatures) )
	stack := make([]int, 0,len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		
	
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack) -1]
			stack = stack[:len(stack) -1]
			output[index] = i - index 
		}
	

		stack = append(stack, i)
	}
	return output
}
