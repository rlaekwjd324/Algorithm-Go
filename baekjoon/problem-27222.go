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
   
   arr := make([]int, n)
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      
      arr[i] = a
   }
   sum := 0
   for i := 0; i < n; i++ {
      var x, y int
      fmt.Fscan(reader, &x, &y)
      
      if arr[i] == 1 && y-x>0 {
         sum += (y-x)
      }
   }
   
   fmt.Fprintf(writer, "%v\n", sum)
}
