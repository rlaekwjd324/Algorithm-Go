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

   var k, n int
   fmt.Fscan(reader, &k, &n)
   
   min := n+1
   max := k*n+1
  
   fmt.Fprintf(writer, "%v %v\n", min, max)
}
