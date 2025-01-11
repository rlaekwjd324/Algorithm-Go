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
   nArr := make([]int, n)
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      nArr[i] = a
   }

   var m int
   fmt.Fscan(reader, &m)
   
   mArr := make([]int, m)
   for i := 0; i < m; i++ {
      var a int
      fmt.Fscan(reader, &a)
      mArr[i] = a
   }

   sum := 0
   for i := 0; i < n; i++ {
      sum += nArr[i]
      if isContainsNumber(sum, mArr) { 
         sum = 0
      }
   }
   
   fmt.Fprintln(writer, sum)
}

func isContainsNumber(i int, arr []int) bool {
   for _, v := range arr {
      if i == v {
         return true
      }
   }
   return false
}
