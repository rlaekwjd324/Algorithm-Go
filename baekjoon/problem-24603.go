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
   var x, y float64
   fmt.Fscan(reader, &n, &x, &y)
   
   for i := 0; i < n; i++ {
      var a float64
      fmt.Fscan(reader, &a)

      fmt.Fprintf(writer, "%v\n", int(a/x*y))
   }
}
