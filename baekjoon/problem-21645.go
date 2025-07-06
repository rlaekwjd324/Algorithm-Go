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

   var n, m, k int
   fmt.Fscan(reader, &n, &m, &k)
   // m = a+bk
   maxA := 0
   maxB := 0
   for i := 1; i <= m; i++ {
      b := i/k
      a := i%k
      if a > k-1 {
         continue
      }
      if maxA < a {
         maxA = a
      }
      if maxB < b {
         maxB = b
      }
   }
   fmt.Fprintf(writer, "%v\n", (maxA+maxB)*n)
}
