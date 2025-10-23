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
   
   var n, p int
   fmt.Fscan(reader, &n, &p)
   
   s := 0
   for i := 0; i < n; i++ {
      s += i+p
   }
   for i := 1; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      
      s -= a
   }

   fmt.Fprintf(writer, "%v\n", s)
}
