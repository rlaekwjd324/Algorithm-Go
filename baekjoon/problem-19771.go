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

   var n int
   fmt.Fscan(reader, &n)
   
   a := 0
   for i := n/2; i > 0; i-- {
      if i*3 > n/2 {
         continue
      }
      a = i*4
      break
   }
   fmt.Fprintf(writer, "%v\n", a)
}
