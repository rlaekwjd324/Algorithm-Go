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
   
   var n, m int
   fmt.Fscan(reader, &n, &m)
   
   if m + 7 == n - 7 {
      fmt.Fprintf(writer, "%v\n", m+7)
      return
   }
   if m + 7 <= 29 {
      fmt.Fprintf(writer, "%v\n", m+7)
      return
   }
   if m + 7 <= 31 && n - 7 <= 0 {
      fmt.Fprintf(writer, "%v\n", m+7)
      return
   }
   fmt.Fprintf(writer, "%v\n", n-7)
}
