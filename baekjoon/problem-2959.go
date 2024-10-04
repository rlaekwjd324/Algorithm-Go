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

   arr := []int{}

   for i := 0; i < 4; i++ {
      var n int
      fmt.Fscan(reader, &n)
      arr = append(arr, n)
   }
   sort.Slice(arr, func(i, j int) bool {
      return arr[i] < arr[j]
   })
   
   fmt.Fprintln(writer, arr[0]*arr[2])
}
