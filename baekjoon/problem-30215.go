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

   var o, a, k int
   fmt.Fscan(reader, &o, &a, &k)
   
   // o = a1 * a + k1 * k
   o -= a
   o -= k
   for i := 0; i <= o/a; i++ {
      for j := 0; j <= o/k; j++ {
         if i * a + j * k == o {
            fmt.Fprintf(writer, "%v\n", 1)
            return
         }
      }
   }
   fmt.Fprintf(writer, "%v\n", 0)
}
