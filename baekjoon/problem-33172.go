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
   
   for i := 0; i < n/2; i++ {
      if n % (i+1) != 0 {
         continue
      }
      isSame := true
      for j := 0; j <= i; j++ {
         t := w[j]
         for k := 1; k < n/(i+1); k++ {
            if w[k*(i+1)+j] != t {
               isSame = false
               break
            }
         }
         if !isSame {
            break
         }
      }
      if isSame {
         fmt.Fprintln(writer, "Yes")
         return
      }
   }
   fmt.Fprintln(writer, "No")
}
