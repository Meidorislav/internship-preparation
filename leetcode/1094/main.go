import "sort"

func carPooling(trips [][]int, capacity int) bool {
    points := make([][]int, 0)
    for _, seg := range trips {
        points = append(points, []int{seg[1], seg[0]})
        points = append(points, []int{seg[2], -seg[0]})
    }

    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] {
            return points[i][1] < points[j][1]
        }
        return points[i][0] < points[j][0]
    })

    curPassengers := 0
    for _, v := range points {
        curPassengers += v[1]
        if curPassengers > capacity {
            return false
        } 
    } 

    return true
}
