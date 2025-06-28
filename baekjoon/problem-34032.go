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

   var n float64
   var s string
   fmt.Fscan(reader, &n, &s)
   
   c := 0
   for _, v := range s {
      if string(v) == "O" {
         c++
      }
   }
   if float64(c) >= n/2.0 {
      fmt.Fprintf(writer, "%v\n", "Yes")
      return
   }
   fmt.Fprintf(writer, "%v\n", "No")
}
