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
   
   for {
      var c, w, l, p float64
      fmt.Fscan(reader, &c, &w, &l, &p)

      if c == 0 && w == 0 && l == 0 && p == 0 {
         return
      }

      count := math.Pow(c, w)
      count = math.Pow(count, l)
      count = math.Pow(count, p)

      fmt.Fprintln(writer, int(count))
   }
}
