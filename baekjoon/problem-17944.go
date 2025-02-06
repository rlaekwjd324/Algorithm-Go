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

   var n, t int
   fmt.Fscan(reader, &n, &t)
   
   h := t % (n * 4 - 2)
   if h > n * 2 {
      h = n*2 - (h - n*2)
   }
   if h == 0 {
      fmt.Fprintln(writer, 2)
      return
   }
   fmt.Fprintln(writer, h)
}
