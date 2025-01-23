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
   var a, b string
   fmt.Fscan(reader, &n, &a, &b)
   
   ac, bc := 0, 0
   for i := 0; i < n; i++ {
      if a[i] == b[i] {
         continue
      }
      if (a[i] == 'R' && b[i] == 'S') || (a[i] == 'P' && b[i] == 'R') || (a[i] == 'S' && b[i] == 'P') {
         ac++
         continue
      }
      bc++
   }

   fmt.Fprintf(writer, "%v %v\n", ac, bc)
}
