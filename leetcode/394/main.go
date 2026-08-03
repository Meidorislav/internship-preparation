func decodeString(s string) string {
    stack_nums := []int{}
    stack_strs := []string{}
    curr_str := ""
    curr_num := 0

    for i := 0; i < len(s); i++ {
        if '0' <= s[i] && s[i] <= '9' {
            curr_num = curr_num * 10 + int(s[i]-'0')
        } else if s[i] == '[' {
            stack_nums = append(stack_nums, curr_num)
            curr_num = 0
			stack_strs = append(stack_strs, curr_str)
			curr_str = ""
        } else if s[i] == ']' {
            count := stack_nums[len(stack_nums)-1]
			prev_str := stack_strs[len(stack_strs)-1]
			repeated := strings.Repeat(curr_str, count)
			curr_str = prev_str + repeated
			stack_strs = stack_strs[:len(stack_strs)-1]
			stack_nums = stack_nums[:len(stack_nums)-1]
        } else {
            curr_str += string(s[i])
        }
    }
    return curr_str
}
