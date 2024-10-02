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

   var n, a, b int
   fmt.Fscan(reader, &n, &a, &b)

   co := 1
   bo := 1

   for i := 0; i < n; i++ {
      co += a
      bo += b
      if co < bo {
         temp := co
         co = bo
         bo = temp
      }
      if co == bo {
         bo--
      }
   }

   fmt.Fprintln(writer, fmt.Sprintf("%v %v", co, bo))
}
