package main

import (
  "bufio"
  "os"
  "strconv"
)

type Point struct {
    r int
    c int
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    buf := make([]byte, 1<<20)
    scanner.Buffer(buf, 1<<20)

    writer := bufio.NewWriterSize(os.Stdout, 1<<20)
    defer writer.Flush()

    readInt := func() int {
        scanner.Scan()
        val, _ := strconv.Atoi(scanner.Text())
        return val
    }

    n := readInt()
    m := readInt()

    matrix := make([][]int, n)
    for i := 0; i < n; i++ {
        matrix[i] = make([]int, m)
        for j := 0; j < m; j++ {
            matrix[i][j] = readInt()
        }
    }
    result := 0

    for i := 0; i <= (n - 1) / 2; i++ {
        for j := 0; j <= (m - 1) / 2; j++ {
            r1 := i
            r2 := n - i - 1
            c1 := j
            c2 := m - j - 1

            points := []Point{{r1, c1}}
            if r1 != r2 {
                points = append(points, Point{r2, c1})
            }
            if c2 != c1 {
                points = append(points, Point{r1, c2})
            }
            if r1 != r2 && c1 != c2 {
                points = append(points, Point{r2, c2})
            }

            counts := make(map[int]int)
            maxFreq := 0
            for _, p := range points {
                val := matrix[p.r][p.c]
                counts[val]++
                if counts[val] > maxFreq {
                    maxFreq = counts[val]
                }
            }
            result += len(points) - maxFreq
        }
    }
    writer.WriteString(strconv.Itoa(result))
    writer.WriteByte('\n')
}
