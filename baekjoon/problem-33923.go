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

   var n, m int
   fmt.Fscan(reader, &n, &m)

   v := 0
   if n == m {
      v = (n-2) * (n-2) + 1
   } else if n > m {
      v = (m-1) * (m-1)
   } else {
      v = (n-1) * (n-1)
   }
   
   fmt.Fprintf(writer, "%v\n", v)
}
