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

   var n, a, b int
   fmt.Fscan(reader, &n, &a, &b)

   count := 0
   for i := 1; i <= n; i++ {
      if i % a != 0 && i % b != 0 {
         continue
      }
      if i % a == 0 && i % b == 0 {
         continue
      }
      count++
   }

   fmt.Fprintln(writer, count)
}
