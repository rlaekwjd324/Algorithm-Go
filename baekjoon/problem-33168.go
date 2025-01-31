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
      var n int
      fmt.Fscan(reader, &n)

      arr[i] = n
   }
   
   for i := 0; i < n-1; i++ {
      t := make([]int, n-i-1)
      for j := 0; j < n-i-1; j++ {
         t[j] = arr[j]+arr[j+1]
         fmt.Fprintf(writer, "%v ", fmt.Sprint(t[j]))
      }
      arr = t
      fmt.Fprintf(writer, "\n")
   }
   
}
