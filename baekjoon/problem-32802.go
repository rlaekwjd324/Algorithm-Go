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

   var n, k int
   fmt.Fscan(reader, &n)

   tArr := make([]string, n)
   sArr := make([]int, n)
   cArr := make([]int, n)
   gArr := make([]int, n)
   for i := 0; i < n; i++ {
      var t string
      var s, c, g int
      fmt.Fscan(reader, &t, &s, &c, &g)

      tArr[i] = t
      sArr[i] = s
      cArr[i] = c
      gArr[i] = g
   }

   fmt.Fscan(reader, &k)

   cSum := 0
   gSum := 0
   for i := 0; i < n; i++ {
      if k >= sArr[i] {
         if tArr[i] == "CAUGHT" {
            cSum += cArr[i]
            gSum += gArr[i]
         } else {
            cSum -= cArr[i]
            gSum -= gArr[i]
         }
      }
   }

   fmt.Fprintf(writer, "%v %v\n", cSum, gSum)
}
