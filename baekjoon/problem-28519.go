package main

import (
   "bufio"
   "fmt"
   "os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   var n, m int
   fmt.Fscan(reader, &n, &m)
   
   if math.Abs(float64(n-m)) <= 2 {
      fmt.Fprintln(writer, n+m)
      return
   }
   if n < m {
      fmt.Fprintln(writer, n*2+1)
      return
   }
   fmt.Fprintln(writer, m*2+1)
}
