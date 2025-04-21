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
   
   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)
      
      c := 0
      sum := 0
      for j := 0; j < n; j++ {
         var a, b int
         fmt.Fscan(reader, &a, &b)

         sum += a
         sum -= b
         if sum < 0 {
            c -= sum
            sum = 0
         }
      }
      fmt.Fprintf(writer, "%v\n", c)
   }
}
