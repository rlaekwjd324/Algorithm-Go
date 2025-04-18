package main

import (
   "bufio"
   "fmt"
   "os"
   "math"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   var n float64
   fmt.Fscan(reader, &n)
   
   s := int(math.Ceil(math.Sqrt(n)))
   
   for i := 0; i < s+2; i++ {
      if i == 0 || i == s+1 {
         fmt.Fprintln(writer, strings.Repeat("x", s+2))
         continue
      }
      fmt.Fprintf(writer, "%s%s%s\n", "x", strings.Repeat(".", s), "x")
   }
}
