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

   switch n % 3 {
   case 1:
      fmt.Fprintln(writer, "U")
   case 2:
      fmt.Fprintln(writer, "O")
   case 0:
      fmt.Fprintln(writer, "S")
   }

}
