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
   
   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var n, x, y int
      fmt.Fscan(reader, &n, &x, &y)
      
      easy := false
      hard := false
      for j := 0; j < n; j++ {
         var a int
         fmt.Fscan(reader, &a)

         if j == 0 && a == x {
            easy = true
         }
         if j == n-1 && a == y {
            hard = true
         }
      }
      if easy && hard {
         fmt.Fprintf(writer, "%v\n", "BOTH")
         continue
      }
      if !easy && !hard {
         fmt.Fprintf(writer, "%v\n", "OKAY")
         continue
      }
      if easy {
         fmt.Fprintf(writer, "%v\n", "EASY")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "HARD")
   }
}
