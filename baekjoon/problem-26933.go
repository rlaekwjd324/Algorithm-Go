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
   
   var n int
   fmt.Fscan(reader, &n)
   
   sum := 0
   for i := 0; i < n; i++ {
      var h, b, k int
      fmt.Fscan(reader, &h, &b, &k)
      
      v := b-h
      if v < 0 {
         v = 0
      }
      sum += v*k
   }
   fmt.Fprintf(writer, "%v\n", sum)
}
