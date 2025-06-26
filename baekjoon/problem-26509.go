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
   
   for i := 0; i < n; i++ {
      var a, b, c, d, e, f int
      fmt.Fscan(reader, &a, &b, &c)
      fmt.Fscan(reader, &d, &e, &f)
      aArr := []int{a, b, c}
      sort.Ints(aArr)
      bArr := []int{d, e, f}
      sort.Ints(bArr)
      if aArr[0]*aArr[0]+aArr[1]*aArr[1] == aArr[2]*aArr[2] && aArr[0] == bArr[0] && aArr[1] == bArr[1] && aArr[2] == bArr[2] {
         fmt.Fprintf(writer, "%v\n", "YES")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "NO")
   }
}
