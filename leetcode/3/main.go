func lengthOfLongestSubstring(s string) int {
    maxCount := 0
    for i := 0; i < len(s); i++ {
        check := make(map[byte]bool)
        count := 0
        
        for j := i; j < len(s); j++ {
            if check[s[j]] {
                break
            }
            check[s[j]] = true
            count++
        }
        
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}

