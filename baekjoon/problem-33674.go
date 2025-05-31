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

   var n, d, k int
   fmt.Fscan(reader, &n, &d, &k)
   
   max := 0
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      if max < a {
         max = a
      }
   }
   c := 0
   sum := 0
   for i := 0; i < d; i++ {
      if sum + max > k {
         sum = 0
         c++
      }
      sum += max
   }
   fmt.Fprintf(writer, "%v\n", c)
}
