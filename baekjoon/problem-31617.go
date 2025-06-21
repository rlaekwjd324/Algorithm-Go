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

   var k, n, m int
   fmt.Fscan(reader, &k, &n)
   
   aArr := make([]int, n)
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      aArr[i] = a
   }
   
   fmt.Fscan(reader, &m)
   bArr := make([]int, m)
   for i := 0; i < m; i++ {
      var b int
      fmt.Fscan(reader, &b)
      bArr[i] = b
   }

   cnt := 0
   for i := 0; i < n; i++ {
      cnt += countOf(aArr[i]+k, bArr)
   }
   
   fmt.Fprintf(writer, "%v\n", cnt)
}

func countOf(t int, arr []int) int {
   cnt := 0
   for _, v := range arr {
      if t == v {
         cnt++
      }
   }
   return cnt
}
