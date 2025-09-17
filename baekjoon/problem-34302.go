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
   
   max := 0
   c := 0
   for i := 0; i < t; i++ {
      var a, b int
      fmt.Fscan(reader, &a, &b)
      if a > b {
         c++
      } else {
         c = 0
      }
      if max < c {
         max = c
      }
   }

   fmt.Fprintf(writer, "%v\n", max)
}
