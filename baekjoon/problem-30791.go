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
   for i := 1; i < 5; i++ {
      var a int
      fmt.Fscan(reader, &a)
      
      if n - a <= 1000 {
         count++
      } else {
         break
      }
   }
   
   fmt.Fprintf(writer, "%v\n", count)
}
