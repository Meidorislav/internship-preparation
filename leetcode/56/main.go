func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    var merged [][]int
    merged = append(merged, intervals[0])

    for i := 1; i < len(intervals); i++ {
        lastIdx := len(merged) - 1
        lastInterval := merged[lastIdx]
        currentInterval := intervals[i]

        if currentInterval[0] <= lastInterval[1] {
            if currentInterval[1] > lastInterval[1] {
                merged[lastIdx][1] = currentInterval[1]
            }
        } else {
            merged = append(merged, currentInterval)
        }
    }

    return merged
}
