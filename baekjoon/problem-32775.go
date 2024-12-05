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

   var s, f int
   fmt.Fscan(reader, &s, &f)

   if s > f {
      fmt.Fprintln(writer, "flight")
      return
   }
   fmt.Fprintln(writer, "high speed rail")
}
