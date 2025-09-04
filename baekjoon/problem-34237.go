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
   
   arr := make([][]int, n)
   for i := 0; i<n; i++ {
      var a, b int
      fmt.Fscan(reader, &a, &b)
      arr[i] = make([]int, 2)
      arr[i][0] = a
      arr[i][1] = b
   }
   gArr := make([][]int, m)
   for i := 0; i<m; i++ {
      var g, x, y int
      fmt.Fscan(reader, &g, &x, &y)
      gArr[i] = make([]int, 3)
      gArr[i][0] = g
      gArr[i][1] = x
      gArr[i][2] = y
   }

   for j := 0; j<m; j++ {
      cnt := 0
      for i := 0; i<n; i++ {
         if gArr[j][1]<=arr[i][0] && gArr[j][2]<=arr[i][1] && arr[i][0]+arr[i][1] <= gArr[j][0] {
            cnt++
         }
      }
      fmt.Fprintf(writer, "%v\n", cnt)
   }
}
