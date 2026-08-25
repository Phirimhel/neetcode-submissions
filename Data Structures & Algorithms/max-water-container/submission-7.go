func maxArea(heights []int) int {

	n := len(heights)
	maxWaterLevel := 0

	i := 0
	j := n-1

	for i < j {

		newWaterLevel := area(heights[i], heights[j], j - i)

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


func area(a , b, containerLength int) int {
	return min(a, b) * containerLength
}

func min(a, b int) int {
	
	if a < b {
		return a 
	}

	if b < a {
		return b
	}

	return a 
}


