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
      
      if d == "ALLOWED" {
         if len(s) != len(p) || !isValid(p, s) {
            fmt.Fprintln(writer, "INTEGRITY OVERFLOW")
            return
         }
      }
      if d == "DENIED"  && s == p {
            fmt.Fprintln(writer, "INTEGRITY OVERFLOW")
            return
      }
   }

   fmt.Fprintln(writer, "SYSTEM SECURE")
}

func isValid(p, s string) bool {
   c := 0
   for i := 0; i < len(p); i++ {
      if p[i] != s[i] {
         c++
      }
      if c > 1 {
         return false
      }
   }
   return true
}
