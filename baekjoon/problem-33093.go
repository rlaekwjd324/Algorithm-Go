package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n, m, k int
   fmt.Fscan(reader, &n, &m, &k)
   
   tArr := []string{}
   a := 0
   aArr := []string{}
   for i := 0; i < n; i++ {
      var s, t string
      fmt.Fscan(reader, &s, &t)
      if isContains(tArr, t) {
         continue
      }
      tArr = append(tArr, t)
      if i < m  {
         continue
      }
      if a < k {
         a++
         aArr = append(aArr, s)
      }
   }

   fmt.Fprintln(writer, a)
   fmt.Fprintln(writer, strings.Join(aArr, "\n"))
}

func isContains(arr []string, s string) bool {
   for _, a := range arr {
      if a == s {
         return true
      }
   }
   return false
}
