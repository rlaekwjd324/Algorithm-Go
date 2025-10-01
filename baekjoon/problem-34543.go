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

   var n, w int
   fmt.Fscan(reader, &n, &w)

   s := n*10
   if n >= 3 {
      s += 20
   }
   if n >= 5 {
      s += 50
   }
   if w > 1000 {
      s -= 15
      if s < 0 {
         s = 0
      }
   }
   
   fmt.Fprintf(writer, "%v\n", s)
}
