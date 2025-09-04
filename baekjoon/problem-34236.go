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

   var n, y1, y2 int
   fmt.Fscan(reader, &n, &y1, &y2)

   y := y1+(y2-y1)*n
  
   fmt.Fprintf(writer, "%v\n", y)
}
