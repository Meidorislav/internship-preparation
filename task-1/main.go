package main 

import "fmt"

func GetExpenses(A, B, C, D int) int {
	if (B >= D) {
		return A
	}
	return A + (D - B) * C 
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(GetExpenses(a, b, c, d))
}
