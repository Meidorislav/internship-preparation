func isOverlapping(a []int, b []int) bool {
    return max(a[0], b[0]) <= min(a[1], b[1])
}

func overlapTwoSegments(a []int, b []int) []int {
    return []int{max(a[0], b[0]), min(a[1], b[1])}
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    result := make([][]int, 0)
    p1 := 0
    p2 := 0

    for p1 < len(firstList) && p2 < len(secondList) {
        if isOverlapping(firstList[p1], secondList[p2]) {
            result = append(result, overlapTwoSegments(firstList[p1], secondList[p2]))
        }

        if firstList[p1][1] < secondList[p2][1] {
            p1++
        } else {
            p2++
        }
    }
    return result
}
