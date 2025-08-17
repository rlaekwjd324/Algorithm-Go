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
   
   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)

      eaten := false
      for j := 0; j < n; j++ {
         var l, r int
         fmt.Fscan(reader, &l, &r)
         if l+r < n {
            eaten = true
         }
      }
      if eaten {
         fmt.Fprintf(writer, "%v\n", "yes")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "no")
   }
}
