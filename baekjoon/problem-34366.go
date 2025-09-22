package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var t int
   fmt.Fscan(reader, &t)
   
   h := 13
   l := 1001
   th := 13
   tl := 8001
   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)

      t := 0
      for j := 0; j < n; j++ {
         var a int
         fmt.Fscan(reader, &a)
         t += a
         if h < a {
            h = a
         }
         if l > a {
            l = a
         }
      }
      if th < t {
         th = t
      }
      if tl > t {
         tl = t
      }
   }
   fmt.Fprintf(writer, "%v\n%v\n%v\n%v\n", h, l, th, tl)
}
