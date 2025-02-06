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

   var n, k int
   var m string
   fmt.Fscan(reader, &n, &m, &k)
   
   if m == "annyong" {
      if k % 2 == 1 {
         fmt.Fprintln(writer, k)
         return
      }
      if k + 1 <= n {
         fmt.Fprintln(writer, k+1)
         return
      }
      fmt.Fprintln(writer, k-1)
      return
   }
   if k % 2 == 0 {
      fmt.Fprintln(writer, k)
      return
   }
   if k + 1 <= n {
      fmt.Fprintln(writer, k+1)
      return
   }
   fmt.Fprintln(writer, k-1)
}
