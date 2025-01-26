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

   var v, w, p, n int
   fmt.Fscan(reader, &v, &w, &p, &n)
   
   if v >= w {
      fmt.Fprintln(writer, "Watermelon")
      return
   }
   if v >= p {
      fmt.Fprintln(writer, "Pomegranates")
      return
   }
   if v >= n {
      fmt.Fprintln(writer, "Nuts")
      return
   }

   fmt.Fprintln(writer, "Nothing")
}
