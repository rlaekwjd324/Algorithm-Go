package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n int
   fmt.Fscan(reader, &n)

   a := ""
   for i := 0; i < n*5; i++ {
      if i < n || i >= 4*n {
         a += strings.Repeat("@", n)
         a += strings.Repeat(" ", n*3)
         a += strings.Repeat("@", n)
      } else if i < n*2 || i >= 3*n {
         a += strings.Repeat("@", n)
         a += strings.Repeat(" ", n*2)
         a += strings.Repeat("@", n)
      } else {
         a += strings.Repeat("@", n*3)
      }
      a += "\n"
   }

   fmt.Fprintln(writer, a)
}
