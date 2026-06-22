package main 

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CountTime(n, t int, stages []int, stage_t int) int {
	first_stage := stages[0] 
	last_stage := stages[n - 1]
	stage_employee_t := stages[stage_t - 1]
	down_diff := stage_employee_t - first_stage
	up_diff := last_stage - stage_employee_t

	if down_diff <= t || up_diff <= t {
		return last_stage - first_stage
	}
	return Min(down_diff, up_diff) + (last_stage - first_stage)
}


func main() {
	var n, t int
	var stage_t int
	fmt.Scan(&n, &t)
	stages := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&stages[i])
	}
	fmt.Scan(&stage_t)

	fmt.Println(CountTime(n, t, stages, stage_t))
}
