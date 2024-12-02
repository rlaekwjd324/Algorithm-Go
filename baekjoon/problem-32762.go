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

   var n, m int
   fmt.Fscan(reader, &n, &m)

   sArr := make([]int, n)
   eArr := make([]int, n)
   for i := 0; i < n; i++ {
      var s, e int
      fmt.Fscan(reader, &s, &e)
      sArr[i] = s
      eArr[i] = e
   }
   sum := 0
   for i := 0; i < m; i++ {
      var t, p int
      fmt.Fscan(reader, &t, &p)

      if isExist(sArr, eArr, t) {
         sum += p
      }
   }
   
   fmt.Fprintf(writer, "%.4f\n", float64(sum)/float64(n))
}

func isExist(s, e []int, t int) bool {
   for i := 0; i < len(s); i++ {
      if t >= s[i] && t <= e[i] {
         return true
      }
   }
   return false
}
