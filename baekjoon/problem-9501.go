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
      var n, d int
      fmt.Fscan(reader, &n, &d)
      
      cnt := 0
      for j := 0; j < n; j++ {
         var v, f, c int
         fmt.Fscan(reader, &v, &f, &c)

         if d/v*c <= f {
            cnt++
         }
      }

      fmt.Fprintf(writer, "%v\n", cnt)
   }
}
