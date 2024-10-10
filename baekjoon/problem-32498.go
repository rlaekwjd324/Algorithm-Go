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

   count := 0
   for i := 0; i < n; i++ {
      var num int
      fmt.Fscanln(reader, &num)

      if num % 2 == 1 {
         count++
      }
   }
   
   fmt.Fprintln(writer, count)
   
}
