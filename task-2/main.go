package main 

import "fmt"
 
func GetSlices(N int) int {
	cuts := 0 
	part := 1 

	for part < N {
		part *= 2 
		cuts++ 
	}
	return cuts 
}

func main() {
	var N int 
	fmt.Scan(&N) 

	fmt.Println(GetSlices(N))
}
