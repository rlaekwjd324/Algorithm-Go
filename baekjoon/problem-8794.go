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

   var z int
   fmt.Fscan(reader, &z)
   
   for i := 0; i < z; i++ {
      c := 0
      var n, m, l int
      fmt.Fscan(reader, &n, &m, &l)
      
      s := n/m
      d := n%m

      if d != 0 && (l == 1 || m-l+1 < d) {
         c++
      }
      c += s
      fmt.Fprintf(writer, "%v\n", c)
   }
}
