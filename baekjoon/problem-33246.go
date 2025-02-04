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
   var p string
   fmt.Fscan(reader, &n, &p)
   
   for i := 0; i < n; i++ {
      var s, d string
      fmt.Fscan(reader, &s, &d)
      
      if d == "ALLOWED" && s != p {
         fmt.Fprintln(writer, "INTEGRITY OVERFLOW")
         return
      }
   }
   fmt.Fprintln(writer, "SYSTEM SECURE")
}
