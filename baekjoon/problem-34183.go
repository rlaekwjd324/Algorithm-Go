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

   var n, m, a, b int
   fmt.Fscan(reader, &n, &m, &a, &b)
   
   if n*3 <= m {
      fmt.Fprintf(writer, "%v\n", 0)
      return
   }
   fmt.Fprintf(writer, "%v\n", (n*3-m)*a+b)
}
