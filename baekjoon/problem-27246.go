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

   var n int64
   fmt.Fscan(reader, &n)

   c := 0
   sum := int64(0)
   for i := int64(1); i <= n; i++ {
      sum += i*i
      if sum > n {
         break
      }
      c++
   }
   
   fmt.Fprintf(writer, "%v\n", c)
}
