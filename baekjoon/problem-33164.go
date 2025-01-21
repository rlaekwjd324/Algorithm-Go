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
   
   nArr := make([]int, n)
   mArr := make([]int, m)
   for i := 0; i < n; i++ {
      var v int
      fmt.Fscan(reader, &v)
      nArr[i] = v
   }
   for i := 0; i < m; i++ {
      var v int
      fmt.Fscan(reader, &v)
      mArr[i] = v
   }
   sum := 0
   for i := 0; i < n; i++ {
      for j := 0; j < m; j++ {
         a := (nArr[i]+mArr[j]) * int(math.Max(float64(nArr[i]), float64(mArr[j])))
         sum += a
      }
   }
   
   fmt.Fprintln(writer, sum)
}
