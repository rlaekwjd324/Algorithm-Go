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
   
   var t int
   fmt.Fscan(reader, &t)

   var arr []int
   for i := 0; i < t; i++ {
      var k int
      fmt.Fscan(reader, &k)

      if len(arr) >= k {
         fmt.Fprintf(writer, "%v\n", arr[k-1])
         continue
      }

      for j := len(arr); j <= k; j++ {
         if j == 0 {
            arr = append(arr, 23)
            continue
         }
         arr = append(arr, arr[j-1]+23)
      }
      fmt.Fprintf(writer, "%v\n", arr[k-1])
   }
}
