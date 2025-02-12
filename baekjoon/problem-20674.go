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
   
   p := 1000
   sum := 0
   for i := 0; i < n; i++ {
      var pi int
      fmt.Fscan(reader, &pi)
      
      if p > pi {
         p = pi
      } else {
         if i > 0 {
            sum += (pi-p)
         }
      }
   }

   fmt.Fprintln(writer, sum)
}
