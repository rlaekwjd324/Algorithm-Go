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

   var n, x float64
   fmt.Fscan(reader, &n, &x)
   
   a := n*(100.0-x)/100
   b := n/a*100-100
  
   fmt.Fprintf(writer, "%.6f\n", b)
}
