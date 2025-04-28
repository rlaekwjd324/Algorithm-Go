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
   
   var t int
   fmt.Fscan(reader, &t)
   
   max := 100
   sum := 100
   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)
      
      sum += n
      if sum > max {
         max = sum
      }
   }
   
   fmt.Fprintf(writer, "%v\n", max)
}
