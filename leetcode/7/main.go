func reverse(x int) int {
    negative := false
    result := 0
    
    if x < 0 {
        negative = true
        x *= (-1)
    }

    for x > 0 {
        result = result * 10 + (x % 10)
        x = x / 10
    }

    if negative {
       result *= (-1) 
    }
    
    if result > math.MaxInt32 || result < math.MinInt32 {
        return 0
    }

    return result
}
