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

   count := 0
   for i := 0; i < 10; i++ {
      var a int
      fmt.Fscan(reader, &a)
      if a == 1 {
         count++
      }
   }
   fmt.Fprintln(writer, count %2)
}
