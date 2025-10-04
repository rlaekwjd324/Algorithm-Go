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

   var x, y, z, u, v, w int
   fmt.Fscan(reader, &x, &y, &z, &u, &v, &w)
   
   s := (u/100) * x + (v/50) * y + (w/20) * z

   fmt.Fprintf(writer, "%v\n", s)
}
