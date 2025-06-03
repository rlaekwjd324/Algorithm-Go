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
   
   max := 1
   c := 1
   cnt := 1

   var pre int
   fmt.Fscan(reader, &pre)

   for i := 1; i < n; i++ {
      var v int
      fmt.Fscan(reader, &v)

      if pre > v {
         cnt++
         if max < c {
            max = c
         }
         c = 1
      } else if i == n-1 {
         c++
         if max < c {
            max = c
         }
      } else {
         c++
      }
      pre = v
   }
   
   fmt.Fprintf(writer, "%v %v\n", cnt, max)
}
