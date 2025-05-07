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
   
   var a, b, c int
   fmt.Fscan(reader, &a, &b, &c)
   
   p := a * b
   g := a + b * c

   if p == g {
      fmt.Fprintf(writer, "%v\n", 0)
      return
   }
   if p > g {
      fmt.Fprintf(writer, "%v\n", 2)
      return
   }
   fmt.Fprintf(writer, "%v\n", 1)
}
