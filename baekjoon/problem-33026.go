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
   var w string
   fmt.Fscan(reader, &n, &w)

   lC := 0
   oC := 0
   for _, s := range w {
      if s == 'L' {
         lC++
      } else {
         oC++
      }
   }
   
   tLc := 0
   tOc := 0
   for i, s := range w {
      if s == 'L' {
         tLc++
      } else {
         tOc++
      }
      if i < len(w)-1 && tLc != lC-tLc && tOc != oC-tOc {
         fmt.Fprintln(writer, i+1)
         return 
      }
   }

   fmt.Fprintln(writer, -1)
}
