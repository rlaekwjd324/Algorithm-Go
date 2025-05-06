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
   
   var a, b, c float64
   fmt.Fscan(reader, &a, &b, &c)
   
   aw := 0.229 * 0.324
   bw := 0.297 * 0.420
   cw := 0.210 * 0.297
   sum := (a * aw * 2.0 + b * bw * 2.0 + c * cw)

   fmt.Fprintf(writer, "%.6f\n", sum)
}
