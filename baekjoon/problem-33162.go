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
   
   sum := 0
   for i := 0; i < n; i++ {
      if i % 2 == 0 {
         sum += 3
      } else {
         sum -= 2
      }
   }

   fmt.Fprintln(writer, sum)
}
