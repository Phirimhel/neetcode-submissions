func maxProfit(prices []int) int {
	l, r := 0, 1 
	max := 0


	for r < len(prices) {
		if prices[l] < prices[r] {
			deal := prices[r] - prices[l]
			if deal > max {
				max = deal
			}
		} else {
			l = r
		}
		r++
		
	}

	return max
}
