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

   for {
      var n int
      _, err := fmt.Fscan(reader, &n)
      if err != nil || n == 0 {
         return
      }

      for i := 0; i < n; i++ {
         fmt.Fprint(writer, strings.Repeat(" ", i))
         fmt.Fprint(writer, "*")
         if i != n-1 {
            fmt.Fprint(writer, strings.Repeat(" ", 2*(n-i)-3))
            fmt.Fprint(writer, "*")
         }
         fmt.Fprint(writer, "\n")
      }
   }
}
