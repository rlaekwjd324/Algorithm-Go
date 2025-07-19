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
   
   f := 0
   min := 1000
   max := 0
   for i := 0; i < n; i++ {
      var x int
      fmt.Fscan(reader, &x)
      if i == 0 {
         f = x
      }
      if min > x {
         min = x
      }
      if max < x {
         max = x
      }
   }
   if f == min {
      fmt.Fprintf(writer, "%v\n", "ez")
      return
   }
   if f == max {
      fmt.Fprintf(writer, "%v\n", "hard")
      return
   }
   fmt.Fprintf(writer, "%v\n", "?")
}
