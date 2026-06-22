package main

import (
	"fmt"
	"sort"
)

func Solution(k int, nums []int) int64 {
	var diffs []int64

	for _, num := range nums {
		pow10 := int64(1)

		for num > 0 {
			digit := num % 10

			diffs = append(diffs, int64(9-digit)*pow10)

			pow10 *= 10
			num /= 10
		}
	}

	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i] > diffs[j]
	})

	var ans int64

	for i := 0; i < k && i < len(diffs); i++ {
		ans += diffs[i]
	}

	return ans
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	nums := make([]int, n)

	for i := range nums {
		fmt.Scan(&nums[i])
	}

	fmt.Println(Solution(k, nums))
}
