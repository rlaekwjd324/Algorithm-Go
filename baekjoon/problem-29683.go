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
   
   var n, a int
   fmt.Fscan(reader, &n, &a)
   
   c := 0
   for i := 0; i < n; i++ {
      var v int
      fmt.Fscan(reader, &v)
      c += (v/a)
   }

   fmt.Fprintf(writer, "%v\n", c)
}
