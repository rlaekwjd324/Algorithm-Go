package main

import (
   "bufio"
   "fmt"
   "os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n int
   fmt.Fscan(reader, &n)
   
   c := 0
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      if a == 1 {
         c++
      }
   }
   v := 0
   if math.Ceil(float64(n)/2.0) > float64(c) {
      v = int(math.Ceil(float64(n)/2.0))-c
   }
   fmt.Fprintf(writer, "%v\n", v)
}
