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

   var w, l int
   fmt.Fscan(reader, &w, &l)
   
   fmt.Fprintf(writer, "%v\n", w*l)
}
