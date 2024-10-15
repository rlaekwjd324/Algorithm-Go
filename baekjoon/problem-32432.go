package main

import (
   "bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n int
   fmt.Fscanln(reader, &n)
   var k int
   fmt.Fscanln(reader, &k)

   a := k-(n-1)
   a = a/n
   
   fmt.Fprintln(writer, a)
}
