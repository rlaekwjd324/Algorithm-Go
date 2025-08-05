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
      var e, n int
      fmt.Fscan(reader, &e, &n)
      
      a := 0
      for j := 0; j < n; j++ {
         var c int
         fmt.Fscan(reader, &c)
         if c > e {
            a++
         }
      }
      fmt.Fprintf(writer, "%v\n", a)
	}
}
