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

   var n, m int64
   fmt.Fscan(reader, &n, &m)

   if n <= 2 || m <= 2 {
      fmt.Fprintf(writer, "%v\n", n*m)
      return
   }
   s := n*m - (n-2)*(m-2)

   fmt.Fprintf(writer, "%v\n", s)
}
