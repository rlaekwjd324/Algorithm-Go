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

   var w int
   fmt.Fscanln(reader, &w)
   var n int
   fmt.Fscanln(reader, &n)

   sum := 0
   for i := 0; i < n; i++ {
      var wi, li int
      fmt.Fscanln(reader, &wi, &li)
      sum += wi*li
   }
   fmt.Fprintln(writer, sum/w)
}
