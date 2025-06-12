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

   var a, b, c float64
   fmt.Fscan(reader, &a, &b, &c)
   
   max := b/c
   max = max*a
   d := a/b
   d = d*c
   if max < d {
      max = d
   }
   
   fmt.Fprintf(writer, "%v\n", int(max))
}
