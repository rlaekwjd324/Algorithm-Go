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

   var m, n, t int
   fmt.Fscan(reader, &m, &n, &t)
   
   fmt.Fprintf(writer, "%v\n", m*n)
}
