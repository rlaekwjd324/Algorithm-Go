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
      var n, a, d int
      fmt.Fscan(reader, &n, &a, &d)
      
      sum := 0
      for  j:=0; j<n; j++ {
         sum += (a+j*d)
      }
      
      fmt.Fprintf(writer, "%v\n", sum)
   }
}
