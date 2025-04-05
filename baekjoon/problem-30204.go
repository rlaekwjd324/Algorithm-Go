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
   
   var n, x int
   fmt.Fscan(reader, &n, &x)
   
   sum := 0
   for i := 0; i < n; i++ {
      var s int
      fmt.Fscan(reader, &s)
      
      sum += s
   }
   if sum % x == 0 {
      fmt.Fprintf(writer, "%v\n", 1)
      return
   }
   fmt.Fprintf(writer, "%v\n", 0)
}
