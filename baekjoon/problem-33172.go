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
      for j := 0; j <= i; j++ {
         isSame := true
         t := w[j]
         for k := 1; k < n/(i+1); k++ {
            fmt.Println(w[j])
            fmt.Println(t)
            if w[k*i+j] != t {
               isSame = false
               break
            }
         }
         if isSame {
            fmt.Fprintln(writer, "Yes")
            return
         }
      }
   }
   fmt.Fprintln(writer, "No")
}
