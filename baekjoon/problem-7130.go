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
   
   var m, h, n int
   fmt.Fscan(reader, &m, &h, &n)
   
   sum := 0
   for i := 0; i < n; i++ {
      var a, b int
      fmt.Fscan(reader, &a, &b)
      
      if a*m > b*h {
         sum += a*m
      } else {
         sum += b*h
      }
   }
   
   fmt.Fprintln(writer, sum)
}
