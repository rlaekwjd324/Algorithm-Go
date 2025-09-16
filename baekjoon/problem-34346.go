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

   var t int
   fmt.Fscan(reader, &t)
   
   if t % 2 == 0 {
      fmt.Fprintf(writer, "%v\n", 2)
      return
   }
   fmt.Fprintf(writer, "%v\n", 1)
}
