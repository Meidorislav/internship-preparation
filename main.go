package main 

import "fmt"

func Solution(l, r int64) int {
	answer := 0 
	for i := 1; i < 10; i++ {
		var num int64
		num = 0 
		for j := 0; j < 18; j++ {
			num = num * int64(10) + int64(i) 
			if num >= l && num <= r {
				answer++
			}
		}
	}
	return answer 
}

func main() {
	var l, r int64
	fmt.Scan(&l, &r)

	fmt.Println(Solution(l, r))
}
