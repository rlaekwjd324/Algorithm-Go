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

   count := 0
   for i := 0; i <= 5; i++ {
      if n-i >5 {
         continue
      }
      if n-i < i {
         break
      }
      count++
   }
   
   fmt.Fprintln(writer, count)
}
