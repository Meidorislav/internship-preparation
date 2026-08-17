package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
)

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)
  writer := bufio.NewWriterSize(os.Stdout, 1<<20)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  m, _ := strconv.Atoi(strings.TrimSpace(line))
  line, _ = reader.ReadString('\n')
  a, _ := strconv.Atoi(strings.TrimSpace(line))
  line, _ = reader.ReadString('\n')
  b, _ := strconv.Atoi(strings.TrimSpace(line))

  ans := (b - a + m) % m

  writer.WriteString(strconv.Itoa(ans))
  writer.WriteByte('\n')
}
