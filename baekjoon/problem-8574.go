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
   
   var n, k, x, y int
   fmt.Fscan(reader, &n, &k, &x, &y)
   
   c := 0
   for i := 0; i < n; i++ {
      var xi, yi int
      fmt.Fscan(reader, &xi, &yi)
      
      sum := (x-xi)*(x-xi)+(y-yi)*(y-yi)
      if sum > k * k {
         c++
      }
   }
   
   fmt.Fprintf(writer, "%v\n", c)
}
