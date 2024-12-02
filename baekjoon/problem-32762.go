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

   for i := 0; i < n; i++ {
      var s, e int
      fmt.Fscan(reader, &s, &e)

   }
   for i := 0; i < m; i++ {
      var t, p int
      fmt.Fscan(reader, &t, &p)

   }
   
   
   fmt.Fprintln(writer, "")
}
