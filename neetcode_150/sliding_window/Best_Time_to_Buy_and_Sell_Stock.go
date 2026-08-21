func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max_profit := 0
	left := 0
	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			left = right
		} else {
			max_profit = Max(max_profit, prices[right]-prices[left])
		}
	}
	return max_profit
}

