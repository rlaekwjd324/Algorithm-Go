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
   p := 0
   for i := 0; i < n; i++ {
      var x int
      fmt.Fscan(reader, &x)

      if i > 0 {
         sum += ((x-p)*(x-p))
      }
      p = x
   }
   fmt.Fprintln(writer, sum)
}
