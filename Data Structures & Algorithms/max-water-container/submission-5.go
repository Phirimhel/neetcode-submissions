func maxArea(heights []int) int {

	n := len(heights)
	maxWaterLevel := 0

	i := 0
	j := n-1

	//[1,7,1,1,1,1,2,5,12,3,500,50,7,8,4,7,38,9,10,12,6]
	// 
	for i < j {

		newWaterLevel := area(heights[i], heights[j], j - i)

		if newWaterLevel > maxWaterLevel {
			maxWaterLevel = newWaterLevel

			// for i < j && area(heights[i+1], heights[j], (j - i)) < maxWaterLevel {
			// 	i++
			// }

			// for i < j && area(heights[i], heights[j-1], (j - i)) < maxWaterLevel {
			// 	j--
			// }
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


