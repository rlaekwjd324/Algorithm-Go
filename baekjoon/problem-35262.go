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

   var n, k int
   var s string
   fmt.Fscan(reader, &n, &k, &s)
  
   c := 0
   for _, v := range s {
      if string(v) == "0" {
         c++
      } else {
         c=0
      }
      if c >= k {
         fmt.Fprintf(writer, "%v\n", 0)
         return
      }
   }
   fmt.Fprintf(writer, "%v\n", 1)
}
