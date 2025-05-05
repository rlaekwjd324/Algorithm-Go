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

   var n,m int
   fmt.Fscan(reader, &n, &m)

   if n < 100 {
      fmt.Fprintln(writer, 1)
      return
   }
   if n*(100-m) < 10000 {
      fmt.Fprintln(writer, 1)
      return
   }
   fmt.Fprintln(writer, 0)
}
