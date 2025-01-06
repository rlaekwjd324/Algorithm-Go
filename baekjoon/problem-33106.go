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

   var n int
   fmt.Fscan(reader, &n)
   for i := 0; i < n; i++ {
      var w string
      fmt.Fscan(reader, &w)
      
      s := ""
      for i := 0; i < len(w); i++ {
         if w[i] == 'c' {
            if i == len(w)-1 || !strings.Contains("ehiy", string(w[i+1])) {
               s += "k"
               continue
            }
            if w[i+1] == 'h' {
               s += "c"
               i++
               continue
            }
            s += "s"
         } else {
            s += string(w[i])
         }
      }

      fmt.Fprintln(writer, s)
   }
}
