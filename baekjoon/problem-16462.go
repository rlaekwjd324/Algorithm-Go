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
   
   sum := 0
   for i := 0; i < n; i++ {
      var s int
      fmt.Fscan(reader, &s)
      
      if s != 100 {
         if s / 10 == 6 {
            s = 90 + s % 10
         }
         if s % 10 == 0 || s % 10 == 6 {
            s = 9 + (s - s % 10)
         }
      }
      sum += s
   }
   a := math.Round(float64(sum) / float64(n))
   fmt.Fprintf(writer, "%v\n", a)
}
