package main

import (
   "bufio"
   "fmt"
   "os"
   "sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   arr := make([]int, 4)
   for i := 0; i < 4; i++ {
      var a int
      fmt.Fscan(reader, &a)
      arr[i] = a
   }
   var b int
   fmt.Fscan(reader, &b)

   sort.Ints(arr)
   if arr[0] == arr[3] {
      fmt.Fprintf(writer, "%v\n", 1)
      return
   }
   if arr[1] != arr[2] {
      fmt.Fprintf(writer, "%v\n", 0)
      return
   }
   if arr[1] == arr[3] && arr[0]+b == arr[3]  {
      fmt.Fprintf(writer, "%v\n", 1)
      return
   }
   
   fmt.Fprintf(writer, "%v\n", 0)
}
