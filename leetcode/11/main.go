func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxArea(height []int) int {
    max_area := 0
    left := 0
    right := len(height) - 1

    for left < right {
        area := min(height[right], height[left]) * (right - left) 
        if area > max_area {
            max_area = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return max_area
}
