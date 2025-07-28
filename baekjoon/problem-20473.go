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

   a3 := n/3
   n -= (a3*3)
   if n%2 == 1 {
      a3--
      n += 3
   }
   a2 := n/2

   fmt.Fprintf(writer, "%v %v\n", a2, a3)
}
