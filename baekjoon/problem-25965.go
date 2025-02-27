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
   
   var n int
   fmt.Fscan(reader, &n)
   
   for i := 0; i < n; i++ {
      var m, k, d, a int
      fmt.Fscan(reader, &m)
      
      var arr []int
      for j := 0; j < m; j++ {
         var kj, dj, aj int
         fmt.Fscan(reader, &kj, &dj, &aj)
         
         arr = append(arr,kj,dj,aj)
      }
      
      fmt.Fscan(reader, &k, &d, &a)

      sum := 0
      for j := 0; j < m; j++ {
         score := arr[j*3]*k-arr[j*3+1]*d+arr[j*3+2]*a
         if score < 0 {
            score = 0
         }
         sum += score
      }

      fmt.Fprintf(writer, "%v\n", sum)
   }
}
