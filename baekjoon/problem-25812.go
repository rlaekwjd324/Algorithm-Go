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

   var t, r int
   fmt.Fscan(reader, &t, &r)
   
   c := 0
   d := 0
   for i := 0; i < t; i++ {
      var v int
      fmt.Fscan(reader, &v)
      
      if v+r > v*2 {
         c++
      } else if v+r < v*2 {
         d++
      }
   }
   if c > d {
      fmt.Fprintf(writer, "%v\n", 1)
      return
   }
   if c < d {
      fmt.Fprintf(writer, "%v\n", 2)
      return
   }
   fmt.Fprintf(writer, "%v\n", 0)
}
