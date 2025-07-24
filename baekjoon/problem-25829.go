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
   
   s1 := 0
   s2 := 0
   se := 0
   for i := 0; i < n; i++ {
      var e, v1, v2 int
      fmt.Fscan(reader, &e, &v1, &v2)
      s1 += v1
      s2 += v2
      if v1 > v2 {
         se += e
      } else {
         se -= e
      }
   }
   if s1 > s2 && se > 0 {
      fmt.Fprintf(writer, "%v\n", 1)
      return
   }
   if s1 < s2 && se < 0 {
      fmt.Fprintf(writer, "%v\n", 2)
      return
   }
   fmt.Fprintf(writer, "%v\n", 0)
}
