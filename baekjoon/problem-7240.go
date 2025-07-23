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

   var n, m int
   fmt.Fscan(reader, &n, &m)
   
   s := 0
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      
      s += a
      if i < n-1 && s > m {
         s--
      }
   }
   fmt.Fprintf(writer, "%v\n", s)
}
