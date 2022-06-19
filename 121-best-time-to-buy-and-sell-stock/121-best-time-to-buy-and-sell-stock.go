
func maxProfit(prices []int) int {
	maxProfit := 0
	currentMin := prices[0]

	// prices = [7,1,5,3,6,4]
	for i := range prices {
		price := prices[i]

		if (maxProfit < (price - currentMin)) {
			maxProfit = price - currentMin
		}

		if price < currentMin {
			currentMin = price
		}
	}

	// output 5
	return maxProfit
}


