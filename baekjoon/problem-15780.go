package main

import (
	"fmt"
	"bufio"
	"os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var n, k int
   fmt.Fscan(reader, &n, &k)
   
   s := 0
   for i := 0; i < k; i++ {
      var a float64
      fmt.Fscan(reader, &a)

      s += int(math.Ceil(a/2.0))
   }
   if s >= n {
      fmt.Fprintf(writer, "%v\n", "YES")
      return
   }
   fmt.Fprintf(writer, "%v\n", "NO")
}
