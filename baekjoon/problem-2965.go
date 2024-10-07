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

   for i := 0; i < 3; i++ {
      var n int
      fmt.Fscan(reader, &n)
      arr = append(arr, n)
   }

   count := 0
   for {
      sort.Slice(arr, func(i, j int) bool {
         return arr[i] < arr[j]
      })
      if arr[1]-arr[0] == 1 && arr[2]-arr[1] == 1 {
         break
      }
      if arr[1]-arr[0] < arr[2]-arr[1] {
         arr[0] = arr[1]+1
         count++
      }else {
         arr[2] = arr[1]-1
         count++
      }
   }
   
   fmt.Fprintln(writer, count)
}
