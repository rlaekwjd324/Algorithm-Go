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

   var s, m, f, b int
   fmt.Fscan(reader, &s)
   fmt.Fscan(reader, &m, &f, &b)

   if s <= 4*60 || m+f+b >= s {
      fmt.Fprintln(writer, "high speed rail")
      return
   }
   fmt.Fprintln(writer, "flight")
}
