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
   
   var a int
   fmt.Fscan(reader, &a)

   y := 2024
   m := 8 + (a-1) * 7
   if m > 12 {
      y += (m/12)
      m = m % 12
   }
   if m == 0 {
      m = 12
      y--
   }

   fmt.Fprintf(writer, "%v %v\n", y, m)
}
