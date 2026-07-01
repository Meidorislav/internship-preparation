package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n+1)
	indeg := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
		indeg[a[i]]++
	}

	o, y := -1, -1
	ok := true
	for v := 1; v <= n; v++ {
		switch indeg[v] {
		case 0:
			if y != -1 {
				ok = false
			}
			y = v
		case 1:
			// норма
		case 2:
			if o != -1 {
				ok = false
			}
			o = v
		default:
			ok = false
		}
	}

	if !ok || o == -1 || y == -1 {
		fmt.Fprintln(writer, -1, -1)
		return
	}

	oneCycle := func() bool {
		cur, cnt := 1, 0
		for {
			cur = a[cur]
			cnt++
			if cnt > n {
				return false
			}
			if cur == 1 {
				break
			}
		}
		return cnt == n
	}

	for x := 1; x <= n; x++ {
		if a[x] != o {
			continue
		}
		old := a[x]
		a[x] = y
		if oneCycle() {
			fmt.Fprintln(writer, x, y)
			return
		}
		a[x] = old
	}

	fmt.Fprintln(writer, -1, -1)
}
