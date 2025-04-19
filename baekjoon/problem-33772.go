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
   var s string
   fmt.Fscan(reader, &n, &s)
   
   a := ""
   for i := 0; i < n; i++ {
      if i+5 > n || s[i:i+5] == "\\../." {
         i += 4
         a += "v"
      } else {
         i += 8
         a += "w"
      }
   }
   fmt.Fprintf(writer, "%v\n", a)
}
