func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func characterReplacement(s string, k int) int {
	counts := [26]int{}

	left := 0
	max_count := 0
	max_length := 0

	for right := 0; right < len(s); right++ {
		counts[s[right]-'A']++
		max_count = Max(max_count, counts[s[right]-'A'])

		for (right-left+1)-max_count > k {
			counts[s[left]-'A']--
			left++
		}

		max_length = Max(max_length, right-left+1)
	}

	return max_length
}

