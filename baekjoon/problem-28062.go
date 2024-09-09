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

   var t int
   fmt.Fscan(reader, &t)

   sum := 0
   min := 1000
   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)

      if min > n && n % 2 ==1 {
         min = n
      }
      sum += n
   }

   if sum % 2 == 1 && min == 1000 {
      fmt.Fprintln(writer, 0)
      return
   }

   if sum % 2 == 1  {
      sum -= min
   }

   fmt.Fprintln(writer, sum)
}
