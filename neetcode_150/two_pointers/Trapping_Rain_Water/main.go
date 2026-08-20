func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left := 0
	right := len(height) - 1
	left_max := height[left]
	right_max := height[right]
	area := 0
	for left < right {
		if left_max <= right_max {
			left++
			if height[left] > left_max {
				left_max = height[left]
			} else {
				area += left_max - height[left]
			}
		} else if right_max < left_max {
			right--
			if height[right] > right_max {
				right_max = height[right]
			} else {
				area += right_max - height[right]
			}
		}
	}
	return area
}

