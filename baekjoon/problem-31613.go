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

   var x, n int
   fmt.Fscan(reader, &x, &n)
   c := 0
   for{
      if x >= n {
         fmt.Fprintf(writer, "%v\n", c)
         return
      }
      r := x % 3
      switch r {
      case 0:
         x++
      case 1:
         x = x*2
      case 2:
         x = x*3
      }
      c++
   }
}
