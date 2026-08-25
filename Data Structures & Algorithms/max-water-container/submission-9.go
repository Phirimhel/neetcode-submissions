func maxArea(heights []int) int {

	n := len(heights)
	maxWaterLevel := 0

	i := 0
	j := n-1

	for i < j {

		newWaterLevel :=  min(heights[i], heights[j]) * (j - i)

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

func min(a, b int) int {
	if a < b {
		return a 
	} 
	return b
}


