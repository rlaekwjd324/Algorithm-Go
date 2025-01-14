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
      var c int
      fmt.Fscan(reader, &c)
      
      arr := make([]int, 8)
      max := 0
      for j := 0; j < c; j++ {
         var x, y int
         fmt.Fscan(reader, &x, &y)

         arr[y-1]++
         if arr[y-1] > max {
            max = arr[y-1]
         }
      }
      fmt.Fprintln(writer, max)
   }

}
