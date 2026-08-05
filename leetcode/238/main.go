func productExceptSelf(nums []int) []int {
    prefix := make([]int, len(nums))
    suffix := make([]int, len(nums))
    prefix[0] = 1
    suffix[len(nums) - 1] = 1
    for i := 1; i < len(nums); i++ {
        prefix[i] = nums[i - 1] * prefix[i - 1]
        suffix[len(nums) - 1 - i] = nums[len(nums) - i] * suffix[len(nums) - i] 
    }
    answer := make([]int, len(nums))
    for i, _ := range answer {
        answer[i] = prefix[i] * suffix[i]
    }

    return answer
}
