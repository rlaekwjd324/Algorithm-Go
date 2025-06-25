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
   max := 0
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      sum += a
      if max < a {
         max = a
      }
   }
   fmt.Fprintf(writer, "%v\n", sum-max)
}
