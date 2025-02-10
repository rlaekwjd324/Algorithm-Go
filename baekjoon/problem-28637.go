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
   
   for i := 0; i < n; i++ {
      var w string
      fmt.Fscan(reader, &w)
      
      a := ""
      for i, v := range w {
         if int(v) >= 65 && int(v) <= 90 {
            if i > 0 {
               a += "_"
            }
            a += string(int(v)+32)
            continue
         }
         a += string(v)
      }
      fmt.Fprintln(writer, a)
   }
}
