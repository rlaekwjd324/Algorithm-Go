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

   var l1, r1, l2, r2, k int
   fmt.Fscan(reader, &l1, &r1, &l2, &r2, &k)
   
   var lm, rm int
   if l1 < l2 {
      lm = l2
   } else {
      lm = l1
   }
   if r1 < r2 {
      rm = r1
   } else {
      rm = r2
   }
   s := 1+rm-lm
   if k <= rm && k >= lm {
      s--
   }
   if s < 0 {
      s = 0
   }
   fmt.Fprintf(writer, "%v\n", s)
}
