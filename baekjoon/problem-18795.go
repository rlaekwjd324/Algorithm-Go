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
   
   sum := 0
   for i := 0; i < n+m; i++ {
      var a int
      fmt.Fscan(reader, &a)
      sum += a
   }
   
   fmt.Fprintf(writer, "%v\n", sum)
}
