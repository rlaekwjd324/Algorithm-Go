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
   
   c := 0
   for i := 0; i < n; i++ {
      var arr []string
      for j := 0; j < m; j++ {
         var s string
         fmt.Fscan(reader, &s)
         
         if !isContains(arr, s) {
            arr = append(arr, s)
         }
      }
      c += (m-len(arr))
   }
   fmt.Fprintf(writer, "%v\n", c)
}

func isContains(arr []string, target string) bool {
   for _, v := range arr {
      if v == target {
         return true
      }
   }
   return false
}
