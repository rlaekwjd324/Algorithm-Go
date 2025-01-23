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
   sum := a+b+c

   if sum <= 21 {
      fmt.Fprintln(writer, 1)
      return
   }

   fmt.Fprintln(writer, 0)
}
