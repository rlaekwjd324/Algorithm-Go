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
   fmt.Fscanln(reader, &n)

   for i := 0; i < n; i++ {
      var a, b int
      fmt.Fscanln(reader, &a, &b)

      s1 := b-a
      if s1 < 0 {
         s1 += 43
      }
      s2 := a-b
      if s2 < 0 {
         s2 += 43
      }
      if s1 < s2 {
         fmt.Fprintln(writer, "Inner circle line")
      } else if s1 == s2 {
         fmt.Fprintln(writer, "Same")
      }else {
         fmt.Fprintln(writer, "Outer circle line")
      }
   }
}
