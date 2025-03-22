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
   
   p := 0
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)

      if p == 0 {
         p = a
      } else {
         if a % p == 0 {
            fmt.Fprintln(writer, a)
            p = 0
         }
      }
   }
}
