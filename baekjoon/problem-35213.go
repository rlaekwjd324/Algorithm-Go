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
  
   fmt.Fprintf(writer, "%v\n", n)
   for i := 0; i<n; i++ {
      fmt.Fprintf(writer, "%v ", 1)
   }
}
