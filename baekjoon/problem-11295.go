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
   
   c := 0
   for {
      var l int
      fmt.Fscan(reader, &l)
      
      if l == 0 {
         break
      }

      c++
      fmt.Fprintf(writer, "User %v\n", c)

      var n int
      fmt.Fscan(reader, &n)
      for i := 0; i < n; i++ {
         var s int
         fmt.Fscan(reader, &s)

         k := float64(s * l) / 100000
         fmt.Fprintf(writer, "%.5f\n", k)
      }
   }
}
