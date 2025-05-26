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
   
   m := 601
   for i := 0; i < n; i++ {
      var a, b, c int
      fmt.Fscan(reader, &a, &b, &c)
      
      s := a + b + c
      if s == 512 {
         fmt.Fprintf(writer, "%v\n", 512)
         return 
      }
      if s > 512 && m > s {
         m = s
      }
   }
   if m == 601 {
      m = -1
   }
   fmt.Fprintf(writer, "%v\n", m)
}
