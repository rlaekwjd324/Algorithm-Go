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

   var r, s, n int
   fmt.Fscan(reader, &r, &s, &n)
   
   if r == 0 {
      fmt.Fprintf(writer, "%v\n", 0)
      return
   }
   sum := 0
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)

      sum += a
      if sum % s == r {
         fmt.Fprintf(writer, "%v\n", i+1)
         return
      }
   }
   fmt.Fprintf(writer, "%v\n", -1)
}
