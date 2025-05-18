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

   var w1, w2, h1, h2 int
   fmt.Fscan(reader, &w1, &h1, &w2, &h2)

   max := w1
   if w1 < w2 {
      max = w2
   }
   sum := (h1+h2+2) * 2 + (max+2) * 2 - 4

   fmt.Fprintf(writer, "%v\n", sum)
}
