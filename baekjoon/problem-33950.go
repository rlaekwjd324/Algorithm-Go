package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   var n int
   fmt.Fscan(reader, &n)

   for i := 0; i < n; i++ {
      var k string
      fmt.Fscan(reader, &k)

      k = strings.ReplaceAll(k, "PO", "PHO")
      fmt.Fprintf(writer, "%v\n", k)
   }
}
