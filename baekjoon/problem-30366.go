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

   var n, m int
   fmt.Fscan(reader, &n, &m)
   
   c := 0
   t := 0
   for i := 0; i < n; i ++ {
      var a int
      fmt.Fscan(reader, &a)
      
      t += a
      if t > m {
         c++
         t = a
      }
      fmt.Fprintf(writer, "%v\n", c)
   }
}
