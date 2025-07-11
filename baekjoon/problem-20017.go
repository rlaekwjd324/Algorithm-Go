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

   var n, m, a int
   fmt.Fscan(reader, &n, &m, &a)
   
   arr := make([][]int, n)
   for i := 0; i < n; i++ {
      temp := make([]int, m)
      for j := 0; j < m; j++ {
         var b int
         fmt.Fscan(reader, &b)
         temp[j] = b
      }
      arr[i] = temp
   }
   c := 0
   for i := 0; i < n-1; i++ {
      for j := 0; j < m; j++ {
         if 2 * arr[i][j] < arr[i+1][j] {
            c++
         }
      }
   }
   fmt.Fprintf(writer, "%v\n", c*a)
}
