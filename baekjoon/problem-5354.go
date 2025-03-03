package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)
      
      for j := 0; j < n; j++ {
         if j == 0 || j == n-1 {
            fmt.Fprintln(writer, strings.Repeat("#", n))
            continue
         }
         
         fmt.Fprint(writer, "#")
         fmt.Fprint(writer, strings.Repeat("J", n-2))
         fmt.Fprint(writer, "#\n")
      }
      
      fmt.Fprint(writer, "\n")
   }
}
