import "slices"

func getHours(speed int, piles []int) int {
    hours := 0
    for _, v := range piles {
        hours += (v + speed - 1) / speed
    }
    return hours
}

func minEatingSpeed(piles []int, h int) int {
    left := 1
    right := slices.Max(piles)
    
    for left < right {
        mid := (left + right) / 2 
        k := getHours(mid, piles)
        if k > h {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return left
}
