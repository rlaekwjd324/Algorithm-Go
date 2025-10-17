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

   var n1, n2 int
   fmt.Fscan(reader, &n1, &n2)
   
   if n1 * 2 <= n2 {
      fmt.Fprintf(writer, "%v\n", "double it")
      return
   }
   fmt.Fprintf(writer, "%v\n", "take it")
}
