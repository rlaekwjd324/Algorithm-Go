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
   
   var m, v, n, k int
   fmt.Fscan(reader, &m, &v, &n, &k)

   fmt.Fprintf(writer, "%v\n", (n+20)*2)
}
