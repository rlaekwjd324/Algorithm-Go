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
      var n string
      var x int
      fmt.Fscan(reader, &n, &x)

      m := 0
      for _, v := range n {
         m = (m*10+int(v-'0')) % x
      }
      
      fmt.Fprintf(writer, "Case %v: %v\n", i+1, m)
   }
}
