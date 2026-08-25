func maxArea(heights []int) int {
	i := 0
	j := len(heights) - 1

	maxWaterLevel := 0
	for i < j {
		
		newWaterLevel := area(heights[i], heights[j], (j - i))
		if newWaterLevel > maxWaterLevel {
			maxWaterLevel = newWaterLevel
		}

		if heights[i] > heights[j] {
			j-- 
		} else {
			i++
		}
	}
	return maxWaterLevel
}

func area(left, right, length int) int {
	return min(left, right) * length
}

func min(a, b int) int {
	if a < b {
		return a 
	} 
	return b
}


