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

   a := n % 8
   switch a {
   case 6:
      a = 4
      break
   case 7:
      a = 3
      break
   case 0:
      a = 2
      break
   }
   
   fmt.Fprintln(writer, a)
}
