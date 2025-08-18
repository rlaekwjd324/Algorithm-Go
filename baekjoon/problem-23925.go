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
      var n, k, s int
      fmt.Fscan(reader, &n, &k, &s)

      l := k-s+n-s
      if l > n {
         l = n
      }
      fmt.Fprintf(writer, "Case #%v: %v\n", i+1, l+k)
   }
}
