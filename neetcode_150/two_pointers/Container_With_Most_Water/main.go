func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	max_area := 0

	for left < right {
		area := Min(heights[left], heights[right]) * (right - left)
		if area > max_area {
			max_area = area
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return max_area
}

