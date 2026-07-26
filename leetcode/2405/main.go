func partitionString(s string) int {
    count := 1
    check := make(map[byte]bool)
    for i := 0; i < len(s); i++ {
        if check[s[i]] {
            count++
            check = make(map[byte]bool)
            i--
        } else {
            check[s[i]] = true
        }
    }
    return count
}
