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
   
   min := 1000000000
   max := -1000000000
   for i := 0; i < n; i++ {
      var x, y int
      fmt.Fscan(reader, &x, &y)
      
      if y < min {
         min = y
      }
      if y > max {
         max = y
      }
   }

   fmt.Fprintln(writer, max-min)
}
