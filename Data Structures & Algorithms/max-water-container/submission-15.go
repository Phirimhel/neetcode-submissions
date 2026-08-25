func maxArea(heights []int) int {
	i := 0
	j := len(heights) - 1

	maxWaterLevel := 0
	for i < j {
		
		if newWaterLevel := min(heights[i], heights[j]) * (j - i); newWaterLevel > maxWaterLevel {
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

func min(a, b int) int {
	if a < b {
		return a 
	} 
	return b
}


