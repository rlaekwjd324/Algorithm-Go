package main

import (
   "bufio"
   "fmt"
   "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n, m int
   fmt.Fscan(reader, &n, &m)

   sum := 0
   var d int
   if n < m {
      sum += n*2
      d = m - n
   } else {
      sum += m*2
      d = n - m
   }
   sum += (d%2)
   fmt.Fprintf(writer, "%v\n", sum)
}
