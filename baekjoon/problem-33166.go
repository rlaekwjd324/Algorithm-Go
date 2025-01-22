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

   var p, q, a, b int
   fmt.Fscan(reader, &p, &q, &a, &b)
   
   if p >= q {
      fmt.Fprintln(writer, q*a)
      return
   }

   sum := (p*a)+(q-p)*b
   fmt.Fprintln(writer, sum)
}
