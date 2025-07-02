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
   
   sum := 0
   arr := []int{}
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      sum += a
      arr = append(arr, a)
   }

   for i := 0; i < n; i++ {
      if sum - arr[i] == arr[i] {
         fmt.Fprintf(writer, "%v\n", i+1)
         return
      }
   }
}
