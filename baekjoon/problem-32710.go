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
   fmt.Fscan(reader, &n)
   
   for i := 1; i <= 9; i++ {
      if n/i <= 9 && n%i == 0 {
         fmt.Fprintln(writer, "1")
         return
      }
   }

   fmt.Fprintln(writer, "0")
}
