package main

import (
	"fmt"
	"bufio"
	"os"
   "slices"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var n, k int
   fmt.Fscan(reader, &n, &k)
   
   var arr []string
   for i := 0; i < n; i++ {
      var d string
      fmt.Fscan(reader, &d)

      if !slices.Contains(arr, d) {
         arr = append(arr, d)
      }
   }

   if len(arr) < k {
      fmt.Fprintf(writer, "%v\n", len(arr))
      return
   }

   fmt.Fprintf(writer, "%v\n", k)
}
