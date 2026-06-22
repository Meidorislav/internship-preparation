package main 

import "fmt"

func Solution(n int, a []int) []int {
	heights := []int{-1, -1}

	for i := 0; i < n; i++ {
		if a[i] % 2 != (i + 1) % 2 {
			if heights[a[i] % 2] == -1 {
				heights[a[i] % 2] = i + 1
			} else {
				return []int{-1, -1}
			}
		}
	}

	if heights[0] == -1 && heights[1] == -1 {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if a[i]%2 == a[j]%2 && (i+1)%2 == (j+1)%2 {
					return []int{i + 1, j + 1}
				}
			}
		}
		return []int{-1, -1}
	}

	if heights[0] == -1 || heights[1] == -1 {
    return []int{-1, -1}
	}
	return heights
}

func main() {
	var n int 
	fmt.Scan(&n)
	a := make([]int, n) 
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	answer := Solution(n, a)
	fmt.Printf("%d %d\n", answer[0], answer[1])
}
