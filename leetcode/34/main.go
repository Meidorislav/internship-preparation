func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    left := 0
    right := len(nums) - 1

    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] >= target {
            right = mid
        } else {
            left = mid + 1
        }
    }

    if nums[left] != target {
        return []int{-1, -1}
    }
    
    start := left

    right = len(nums) - 1
    for left < right {
        mid := left + (right - left + 1) / 2
        
        if nums[mid] <= target {
            left = mid
        } else {
            right = mid - 1
        }
    }

    return []int{start, left}
}
