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

   var n, m, k int
   var s string
   fmt.Fscan(reader, &n, &m, &k, &s)
   
   arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
   v := arr[(n-1)*2+m-1]
   for i := 0; i < k; i++ {
      var temp []int
      switch string(s[i]) {
      case "A":
         temp = []int{arr[4], arr[5], arr[6], arr[7], arr[0], arr[1], arr[2], arr[3]}
      case "B":
         temp = []int{arr[3], arr[2], arr[1], arr[0], arr[7], arr[6], arr[5], arr[4]}
      case "C":
         temp = []int{arr[7], arr[6], arr[5], arr[4], arr[3], arr[2], arr[1], arr[0]}
      case "D":
         temp = []int{arr[2], arr[0], arr[4], arr[1], arr[6], arr[3], arr[7], arr[5]}
      }
      arr = temp
   }
   i := indexOf(arr, v)
   // TODO: if i == -1 -> error
   fmt.Fprintf(writer, "%v %v\n", i/2+1, i%2+1)
}
func indexOf(arr []int, target int) int {
   for i, v := range arr {
      if v == target {
         return i
      }
   }
   return -1
}
