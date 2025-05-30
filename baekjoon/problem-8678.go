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
   
   for i := 0; i < n; i++ {
      var a, b int
      fmt.Fscan(reader, &a, &b)
      if b%a == 0 {
         fmt.Fprintf(writer, "%s\n", "TAK")
         continue
      }
      fmt.Fprintf(writer, "%s\n", "NIE")
   }
   
}
