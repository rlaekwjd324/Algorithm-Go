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
   
   min := 2147483648
   for i := 0; i < n; i++ {
      var m, o int
      fmt.Fscan(reader, &m, &o)

      if o == 0 && min > m {
         min = m
      }
   }
   
   if min == 2147483648 {
      fmt.Fprintln(writer, -1)
      return
   }
   fmt.Fprintln(writer, min)
}
