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
   
   arr := make([]int, n)
   max := 0
   for i := 0; i < 2*n; i++ {
      var x int
      fmt.Fscan(reader, &x)

      if arr[x-1] == 0 {
         arr[x-1] = i+1
      } else {
         t := i - arr[x-1]
         if max < t {
            max = t
         }
      }
   }

   fmt.Fprintf(writer, "%v\n", max)
}
