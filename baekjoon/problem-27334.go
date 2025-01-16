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

   var n int
   fmt.Fscan(reader, &n)
   
   arr := make([]int, n)
   new := make([]int, n)
   for i := 0; i < n; i++ {
      var n int
      fmt.Fscan(reader, &n)
      
      arr[i] = n
      new[i] = n
   }
   sort.Slice(arr, func(a, b int) bool {
      return arr[a] < arr[b]
   })
   for i := 0; i < n; i++ {
      index := indexOf(arr, new[i])+1
      fmt.Fprintln(writer, index)
   }
}

func indexOf(arr []int, target int) int {
   for i, v := range arr {
      if v == target {
         return i
      }
   }
   return -1
}
