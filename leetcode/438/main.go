func findAnagrams(s string, p string) []int {
    result := []int{}

    if len(s) < len(p) {
        return result
    }
    
    var sCount [26]int
    var pCount [26]int

    for i := 0; i < len(p); i++ {
        sCount[s[i] - 'a']++
        pCount[p[i] - 'a']++
    }

    if pCount == sCount { 
        result = append(result, 0)
    }

    start := 0
    for end := len(p); end < len(s); end++ {
        sCount[s[end]-'a']++
        sCount[s[start]-'a']--
        start++

        if pCount == sCount {
            result = append(result, start)
        }
    }

    return result
}
