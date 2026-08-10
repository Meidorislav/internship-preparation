func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func partition(s string) [][]string {
    var result [][]string
	var current []string

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			temp := make([]string, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]

			if isPalindrome(sub) {
				current = append(current, sub) 
				backtrack(end)
				current = current[:len(current)-1]
			}
		}
	}

	backtrack(0)
	return result
}
