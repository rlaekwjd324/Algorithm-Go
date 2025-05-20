package main

import (
   "bufio"
   "fmt"
   "os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var w, h, n, a, b float64
   fmt.Fscan(reader, &w, &h, &n, &a, &b)

   if w < a || h < b {
      fmt.Fprintf(writer, "%v\n", -1)
      return
   }
   sum := math.Ceil(n / (math.Floor(w/a) * math.Floor(h/b)))
   fmt.Fprintf(writer, "%v\n", sum)
}
