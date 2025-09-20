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

   var n float64
   fmt.Fscan(reader, &n)
 
   fmt.Fprintf(writer, "%v\n", n-0.3)
}
