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

   fmt.Fprintf(writer, "%v\n", getSch(arr))
}

func getSch(arr []int) string {
   for _, v := range arr {
      if v == 3 {
         return "None"
      }
   }
   isSu := true
   for _, v := range arr {
      if v != 5 {
         isSu = false
         break
      }
   }
   if isSu {
      return "Named"
   }
   sum := 0
   for _, v := range arr {
      sum += v
   }
   if float64(float64(sum)/float64(len(arr))) >= 4.5 {
      return "High"
   }
   return "Common"
}
