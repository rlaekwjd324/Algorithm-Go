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
   
   c := 0
   for i := 0; i < n; i++ {
      var s string
      fmt.Fscan(reader, &s)
      
      if string(s[0]) == "C" {
         c++
      }
   }

   fmt.Fprintf(writer, "%v\n", c)
}
