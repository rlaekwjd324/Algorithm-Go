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
   var t, s float64
   fmt.Fscan(reader, &t, &n, &s)
   
   for i := 0; i < n; i++ {
      var f, b float64
      fmt.Fscan(reader, &f, &b)
      
      if t/f+t/b < s {
         fmt.Fprintln(writer, "HOPE")
         return
      }
   }
   fmt.Fprintln(writer, "DOOMED")

}
