func isPalindrome(s string) bool {
	filtred := []rune{}
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') ||
			(c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') {
			filtred = append(filtred, unicode.ToLower(c))
		}
	}

	left := 0
	right := len(filtred) - 1

	for right > left {
		if filtred[left] != filtred[right] {
			return false
		}
		left++
		right--
	}
	return true
}

