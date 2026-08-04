func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    stack := make([]int, 0)
    result := make([]int, n)

    for i := n - 1; i >= 0; i-- {
        for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) > 0 {
            result[i] = stack[len(stack)-1] - i
        }
        stack = append(stack, i)
    }

    return result
}
