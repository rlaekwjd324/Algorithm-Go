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

   var h, l, a, b int
   fmt.Fscan(reader, &h, &l, &a, &b)
   
	if h >= (a + 1) / 2 && l >= b {
      fmt.Fprintf(writer, "%v\n", "YES")
   } else if h >= (b + 1) / 2 && l >= a {
      fmt.Fprintf(writer, "%v\n", "YES")
   } else {
      fmt.Fprintf(writer, "%v\n", "NO")
   }
}
