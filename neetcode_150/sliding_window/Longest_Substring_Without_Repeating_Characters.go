func lengthOfLongestSubstring(s string) int {
	chars := make(map[byte]int)

	left := 0
	max_len := 0

	for right := 0; right < len(s); right++ {
		if prev, ok := chars[s[right]]; ok && prev >= left {
			left = prev + 1
		}

		chars[s[right]] = right

		length := right - left + 1
		if length > max_len {
			max_len = length
		}
	}

	return max_len
}
